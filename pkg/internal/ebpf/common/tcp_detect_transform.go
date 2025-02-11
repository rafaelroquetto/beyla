package ebpfcommon

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/cilium/ebpf/ringbuf"

	"github.com/grafana/beyla/pkg/config"
	"github.com/grafana/beyla/pkg/internal/request"
)

// nolint:cyclop
func ReadTCPRequestIntoSpan(cfg *config.EBPFTracer, record *ringbuf.Record, filter ServiceFilter) (request.Span, bool, error) {
	var event TCPRequestInfo

	err := binary.Read(bytes.NewBuffer(record.RawSample), binary.LittleEndian, &event)
	if err != nil {
		return request.Span{}, true, err
	}

	if !filter.ValidPID(event.Pid.UserPid, event.Pid.Ns, PIDTypeKProbes) {
		return request.Span{}, true, nil
	}

	l := int(event.Len)
	if l < 0 || len(event.Buf) < l {
		l = len(event.Buf)
	}

	rl := int(event.RespLen)
	if rl < 0 || len(event.Rbuf) < rl {
		rl = len(event.Rbuf)
	}

	b := event.Buf[:l]

	if cfg.ProtocolDebug {
		fmt.Printf("[>] %v\n", b)
		fmt.Printf("[<] %v\n", event.Rbuf[:rl])
	}

	// Check if we have a SQL statement
	op, table, sql, kind := detectSQLPayload(cfg.HeuristicSQLDetect, b)
	if validSQL(op, table) {
		return TCPToSQLToSpan(&event, op, table, sql, kind), false, nil
	} else {
		op, table, sql, kind = detectSQLPayload(cfg.HeuristicSQLDetect, event.Rbuf[:rl])
		if validSQL(op, table) {
			reverseTCPEvent(&event)

			return TCPToSQLToSpan(&event, op, table, sql, kind), false, nil
		}
	}

	if maybeFastCGI(b) {
		op, uri, status := detectFastCGI(b, event.Rbuf[:rl])
		if status >= 0 {
			return TCPToFastCGIToSpan(&event, op, uri, status), false, nil
		}
	}

	switch {
	case isRedis(b) && isRedis(event.Rbuf[:rl]):
		op, text, ok := parseRedisRequest(string(b))

		if ok {
			var status int
			if op == "" {
				op, text, ok = parseRedisRequest(string(event.Rbuf[:rl]))
				if !ok || op == "" {
					return request.Span{}, true, nil // ignore if we couldn't parse it
				}
				// We've caught the event reversed in the middle of communication, let's
				// reverse the event
				reverseTCPEvent(&event)
				status = redisStatus(b)
			} else {
				status = redisStatus(event.Rbuf[:rl])
			}

			return TCPToRedisToSpan(&event, op, text, status), false, nil
		}
	default:
		// Kafka and gRPC can look very similar in terms of bytes. We can mistake one for another.
		// We try gRPC first because it's more reliable in detecting false gRPC sequences.
		if isHTTP2(b, int(event.Len)) || isHTTP2(event.Rbuf[:rl], int(event.RespLen)) {
			MisclassifiedEvents <- MisclassifiedEvent{EventType: EventTypeKHTTP2, TCPInfo: &event}
		} else {
			k, err := ProcessPossibleKafkaEvent(&event, b, event.Rbuf[:rl])
			if err == nil {
				return TCPToKafkaToSpan(&event, k), false, nil
			}
		}
	}

	return request.Span{}, true, nil // ignore if we couldn't parse it
}

func reverseTCPEvent(trace *TCPRequestInfo) {
	if trace.Direction == 0 {
		trace.Direction = 1
	} else {
		trace.Direction = 0
	}

	port := trace.ConnInfo.S_port
	addr := trace.ConnInfo.S_addr
	trace.ConnInfo.S_addr = trace.ConnInfo.D_addr
	trace.ConnInfo.S_port = trace.ConnInfo.D_port
	trace.ConnInfo.D_addr = addr
	trace.ConnInfo.D_port = port
}
