// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package gotracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_tpConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpf_tpEgressKeyT struct {
	S_port uint16
	D_port uint16
}

type bpf_tpFramerFuncInvocationT struct {
	FramerPtr uint64
	Tp        bpf_tpTpInfoT
	InitialN  int64
}

type bpf_tpGoAddrKeyT struct {
	Pid  uint64
	Addr uint64
}

type bpf_tpGoroutineMetadata struct {
	Parent    bpf_tpGoAddrKeyT
	Timestamp uint64
}

type bpf_tpGrpcClientFuncInvocationT struct {
	StartMonotimeNs uint64
	Cc              uint64
	Method          uint64
	MethodLen       uint64
	Tp              bpf_tpTpInfoT
	Flags           uint64
}

type bpf_tpGrpcFramerFuncInvocationT struct {
	FramerPtr uint64
	Tp        bpf_tpTpInfoT
	Offset    int64
}

type bpf_tpGrpcSrvFuncInvocationT struct {
	StartMonotimeNs uint64
	Stream          uint64
	Tp              bpf_tpTpInfoT
}

type bpf_tpGrpcTransportsT struct {
	Type uint8
	_    [3]byte
	Conn bpf_tpConnectionInfoT
}

type bpf_tpHttpClientDataT struct {
	Method        [7]uint8
	Path          [100]uint8
	_             [5]byte
	ContentLength int64
	Pid           struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	_ [4]byte
}

type bpf_tpHttpFuncInvocationT struct {
	StartMonotimeNs uint64
	Tp              bpf_tpTpInfoT
}

type bpf_tpKafkaClientReqT struct {
	Type            uint8
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [256]uint8
	_               [7]byte
	Conn            bpf_tpConnectionInfoT
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
}

type bpf_tpKafkaGoReqT struct {
	Type            uint8
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Topic           [64]uint8
	_               [7]byte
	Conn            bpf_tpConnectionInfoT
	Tp              bpf_tpTpInfoT
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Op uint8
	_  [7]byte
}

type bpf_tpNewFuncInvocationT struct{ Parent uint64 }

type bpf_tpOffTableT struct{ Table [43]uint64 }

type bpf_tpProduceReqT struct {
	MsgPtr          uint64
	ConnPtr         uint64
	StartMonotimeNs uint64
}

type bpf_tpRedisClientReqT struct {
	Type            uint8
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [256]uint8
	_               [7]byte
	Conn            bpf_tpConnectionInfoT
	_               [4]byte
	Tp              bpf_tpTpInfoT
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Err uint8
	_   [3]byte
}

type bpf_tpServerHttpFuncInvocationT struct {
	StartMonotimeNs uint64
	Tp              bpf_tpTpInfoT
	Method          [7]uint8
	Path            [100]uint8
	_               [5]byte
	ContentLength   uint64
	Status          uint64
}

type bpf_tpSqlFuncInvocationT struct {
	StartMonotimeNs uint64
	SqlParam        uint64
	QueryLen        uint64
	Conn            bpf_tpConnectionInfoT
	_               [4]byte
	Tp              bpf_tpTpInfoT
}

type bpf_tpTopicT struct {
	Name [64]int8
	Tp   bpf_tpTpInfoT
}

type bpf_tpTpInfoPidT struct {
	Tp    bpf_tpTpInfoT
	Pid   uint32
	Valid uint8
	_     [3]byte
}

type bpf_tpTpInfoT struct {
	TraceId  [16]uint8
	SpanId   [8]uint8
	ParentId [8]uint8
	Ts       uint64
	Flags    uint8
	_        [7]byte
}

// loadBpf_tp returns the embedded CollectionSpec for bpf_tp.
func loadBpf_tp() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_tpBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_tp: %w", err)
	}

	return spec, err
}

// loadBpf_tpObjects loads bpf_tp and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_tpObjects
//	*bpf_tpPrograms
//	*bpf_tpMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_tpObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_tp()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_tpSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpSpecs struct {
	bpf_tpProgramSpecs
	bpf_tpMapSpecs
}

// bpf_tpSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpProgramSpecs struct {
	UprobeClientConnClose                     *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_Close"`
	UprobeClientConnInvoke                    *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_Invoke"`
	UprobeClientConnInvokeReturn              *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_Invoke_return"`
	UprobeClientConnNewStream                 *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_NewStream"`
	UprobeClientConnNewStreamReturn           *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_NewStream_return"`
	UprobeServeHTTP                           *ebpf.ProgramSpec `ebpf:"uprobe_ServeHTTP"`
	UprobeServeHTTPReturns                    *ebpf.ProgramSpec `ebpf:"uprobe_ServeHTTPReturns"`
	UprobeClientStreamRecvMsgReturn           *ebpf.ProgramSpec `ebpf:"uprobe_clientStream_RecvMsg_return"`
	UprobeClientRoundTrip                     *ebpf.ProgramSpec `ebpf:"uprobe_client_roundTrip"`
	UprobeConnServe                           *ebpf.ProgramSpec `ebpf:"uprobe_connServe"`
	UprobeConnServeRet                        *ebpf.ProgramSpec `ebpf:"uprobe_connServeRet"`
	UprobeExecDC                              *ebpf.ProgramSpec `ebpf:"uprobe_execDC"`
	UprobeGrpcFramerWriteHeaders              *ebpf.ProgramSpec `ebpf:"uprobe_grpcFramerWriteHeaders"`
	UprobeGrpcFramerWriteHeadersReturns       *ebpf.ProgramSpec `ebpf:"uprobe_grpcFramerWriteHeaders_returns"`
	UprobeHttp2FramerWriteHeaders             *ebpf.ProgramSpec `ebpf:"uprobe_http2FramerWriteHeaders"`
	UprobeHttp2FramerWriteHeadersReturns      *ebpf.ProgramSpec `ebpf:"uprobe_http2FramerWriteHeaders_returns"`
	UprobeHttp2ResponseWriterStateWriteHeader *ebpf.ProgramSpec `ebpf:"uprobe_http2ResponseWriterStateWriteHeader"`
	UprobeHttp2RoundTrip                      *ebpf.ProgramSpec `ebpf:"uprobe_http2RoundTrip"`
	UprobeHttp2ServerOperateHeaders           *ebpf.ProgramSpec `ebpf:"uprobe_http2Server_operateHeaders"`
	UprobeHttp2serverConnRunHandler           *ebpf.ProgramSpec `ebpf:"uprobe_http2serverConn_runHandler"`
	UprobeNetFdRead                           *ebpf.ProgramSpec `ebpf:"uprobe_netFdRead"`
	UprobePersistConnRoundTrip                *ebpf.ProgramSpec `ebpf:"uprobe_persistConnRoundTrip"`
	UprobeProcGoexit1                         *ebpf.ProgramSpec `ebpf:"uprobe_proc_goexit1"`
	UprobeProcNewproc1                        *ebpf.ProgramSpec `ebpf:"uprobe_proc_newproc1"`
	UprobeProcNewproc1Ret                     *ebpf.ProgramSpec `ebpf:"uprobe_proc_newproc1_ret"`
	UprobeProtocolRoundtrip                   *ebpf.ProgramSpec `ebpf:"uprobe_protocol_roundtrip"`
	UprobeProtocolRoundtripRet                *ebpf.ProgramSpec `ebpf:"uprobe_protocol_roundtrip_ret"`
	UprobeQueryDC                             *ebpf.ProgramSpec `ebpf:"uprobe_queryDC"`
	UprobeQueryReturn                         *ebpf.ProgramSpec `ebpf:"uprobe_queryReturn"`
	UprobeReadRequestReturns                  *ebpf.ProgramSpec `ebpf:"uprobe_readRequestReturns"`
	UprobeReadRequestStart                    *ebpf.ProgramSpec `ebpf:"uprobe_readRequestStart"`
	UprobeReaderRead                          *ebpf.ProgramSpec `ebpf:"uprobe_reader_read"`
	UprobeReaderReadRet                       *ebpf.ProgramSpec `ebpf:"uprobe_reader_read_ret"`
	UprobeReaderSendMessage                   *ebpf.ProgramSpec `ebpf:"uprobe_reader_send_message"`
	UprobeRedisProcess                        *ebpf.ProgramSpec `ebpf:"uprobe_redis_process"`
	UprobeRedisProcessRet                     *ebpf.ProgramSpec `ebpf:"uprobe_redis_process_ret"`
	UprobeRedisWithWriter                     *ebpf.ProgramSpec `ebpf:"uprobe_redis_with_writer"`
	UprobeRedisWithWriterRet                  *ebpf.ProgramSpec `ebpf:"uprobe_redis_with_writer_ret"`
	UprobeRoundTrip                           *ebpf.ProgramSpec `ebpf:"uprobe_roundTrip"`
	UprobeRoundTripReturn                     *ebpf.ProgramSpec `ebpf:"uprobe_roundTripReturn"`
	UprobeSaramaBrokerWrite                   *ebpf.ProgramSpec `ebpf:"uprobe_sarama_broker_write"`
	UprobeSaramaResponsePromiseHandle         *ebpf.ProgramSpec `ebpf:"uprobe_sarama_response_promise_handle"`
	UprobeSaramaSendInternal                  *ebpf.ProgramSpec `ebpf:"uprobe_sarama_sendInternal"`
	UprobeServerHandleStream                  *ebpf.ProgramSpec `ebpf:"uprobe_server_handleStream"`
	UprobeServerHandleStreamReturn            *ebpf.ProgramSpec `ebpf:"uprobe_server_handleStream_return"`
	UprobeServerHandlerTransportHandleStreams *ebpf.ProgramSpec `ebpf:"uprobe_server_handler_transport_handle_streams"`
	UprobeTransportHttp2ClientNewStream       *ebpf.ProgramSpec `ebpf:"uprobe_transport_http2Client_NewStream"`
	UprobeTransportWriteStatus                *ebpf.ProgramSpec `ebpf:"uprobe_transport_writeStatus"`
	UprobeWriteSubset                         *ebpf.ProgramSpec `ebpf:"uprobe_writeSubset"`
	UprobeWriterProduce                       *ebpf.ProgramSpec `ebpf:"uprobe_writer_produce"`
	UprobeWriterWriteMessages                 *ebpf.ProgramSpec `ebpf:"uprobe_writer_write_messages"`
}

// bpf_tpMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpMapSpecs struct {
	Events                        *ebpf.MapSpec `ebpf:"events"`
	FetchRequests                 *ebpf.MapSpec `ebpf:"fetch_requests"`
	FramerInvocationMap           *ebpf.MapSpec `ebpf:"framer_invocation_map"`
	GoOffsetsMap                  *ebpf.MapSpec `ebpf:"go_offsets_map"`
	GoTraceMap                    *ebpf.MapSpec `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap     *ebpf.MapSpec `ebpf:"golang_mapbucket_storage_map"`
	GrpcFramerInvocationMap       *ebpf.MapSpec `ebpf:"grpc_framer_invocation_map"`
	HeaderReqMap                  *ebpf.MapSpec `ebpf:"header_req_map"`
	Http2ReqMap                   *ebpf.MapSpec `ebpf:"http2_req_map"`
	IncomingTraceMap              *ebpf.MapSpec `ebpf:"incoming_trace_map"`
	KafkaRequests                 *ebpf.MapSpec `ebpf:"kafka_requests"`
	Newproc1                      *ebpf.MapSpec `ebpf:"newproc1"`
	OngoingClientConnections      *ebpf.MapSpec `ebpf:"ongoing_client_connections"`
	OngoingGoHttp                 *ebpf.MapSpec `ebpf:"ongoing_go_http"`
	OngoingGoroutines             *ebpf.MapSpec `ebpf:"ongoing_goroutines"`
	OngoingGrpcClientRequests     *ebpf.MapSpec `ebpf:"ongoing_grpc_client_requests"`
	OngoingGrpcHeaderWrites       *ebpf.MapSpec `ebpf:"ongoing_grpc_header_writes"`
	OngoingGrpcOperateHeaders     *ebpf.MapSpec `ebpf:"ongoing_grpc_operate_headers"`
	OngoingGrpcRequestStatus      *ebpf.MapSpec `ebpf:"ongoing_grpc_request_status"`
	OngoingGrpcServerRequests     *ebpf.MapSpec `ebpf:"ongoing_grpc_server_requests"`
	OngoingGrpcTransports         *ebpf.MapSpec `ebpf:"ongoing_grpc_transports"`
	OngoingHttpClientRequests     *ebpf.MapSpec `ebpf:"ongoing_http_client_requests"`
	OngoingHttpClientRequestsData *ebpf.MapSpec `ebpf:"ongoing_http_client_requests_data"`
	OngoingHttpServerRequests     *ebpf.MapSpec `ebpf:"ongoing_http_server_requests"`
	OngoingKafkaRequests          *ebpf.MapSpec `ebpf:"ongoing_kafka_requests"`
	OngoingProduceMessages        *ebpf.MapSpec `ebpf:"ongoing_produce_messages"`
	OngoingProduceTopics          *ebpf.MapSpec `ebpf:"ongoing_produce_topics"`
	OngoingRedisRequests          *ebpf.MapSpec `ebpf:"ongoing_redis_requests"`
	OngoingServerConnections      *ebpf.MapSpec `ebpf:"ongoing_server_connections"`
	OngoingSqlQueries             *ebpf.MapSpec `ebpf:"ongoing_sql_queries"`
	OngoingStreams                *ebpf.MapSpec `ebpf:"ongoing_streams"`
	OutgoingTraceMap              *ebpf.MapSpec `ebpf:"outgoing_trace_map"`
	ProduceRequests               *ebpf.MapSpec `ebpf:"produce_requests"`
	ProduceTraceparents           *ebpf.MapSpec `ebpf:"produce_traceparents"`
	RedisWrites                   *ebpf.MapSpec `ebpf:"redis_writes"`
	TraceMap                      *ebpf.MapSpec `ebpf:"trace_map"`
}

// bpf_tpObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpObjects struct {
	bpf_tpPrograms
	bpf_tpMaps
}

func (o *bpf_tpObjects) Close() error {
	return _Bpf_tpClose(
		&o.bpf_tpPrograms,
		&o.bpf_tpMaps,
	)
}

// bpf_tpMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpMaps struct {
	Events                        *ebpf.Map `ebpf:"events"`
	FetchRequests                 *ebpf.Map `ebpf:"fetch_requests"`
	FramerInvocationMap           *ebpf.Map `ebpf:"framer_invocation_map"`
	GoOffsetsMap                  *ebpf.Map `ebpf:"go_offsets_map"`
	GoTraceMap                    *ebpf.Map `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap     *ebpf.Map `ebpf:"golang_mapbucket_storage_map"`
	GrpcFramerInvocationMap       *ebpf.Map `ebpf:"grpc_framer_invocation_map"`
	HeaderReqMap                  *ebpf.Map `ebpf:"header_req_map"`
	Http2ReqMap                   *ebpf.Map `ebpf:"http2_req_map"`
	IncomingTraceMap              *ebpf.Map `ebpf:"incoming_trace_map"`
	KafkaRequests                 *ebpf.Map `ebpf:"kafka_requests"`
	Newproc1                      *ebpf.Map `ebpf:"newproc1"`
	OngoingClientConnections      *ebpf.Map `ebpf:"ongoing_client_connections"`
	OngoingGoHttp                 *ebpf.Map `ebpf:"ongoing_go_http"`
	OngoingGoroutines             *ebpf.Map `ebpf:"ongoing_goroutines"`
	OngoingGrpcClientRequests     *ebpf.Map `ebpf:"ongoing_grpc_client_requests"`
	OngoingGrpcHeaderWrites       *ebpf.Map `ebpf:"ongoing_grpc_header_writes"`
	OngoingGrpcOperateHeaders     *ebpf.Map `ebpf:"ongoing_grpc_operate_headers"`
	OngoingGrpcRequestStatus      *ebpf.Map `ebpf:"ongoing_grpc_request_status"`
	OngoingGrpcServerRequests     *ebpf.Map `ebpf:"ongoing_grpc_server_requests"`
	OngoingGrpcTransports         *ebpf.Map `ebpf:"ongoing_grpc_transports"`
	OngoingHttpClientRequests     *ebpf.Map `ebpf:"ongoing_http_client_requests"`
	OngoingHttpClientRequestsData *ebpf.Map `ebpf:"ongoing_http_client_requests_data"`
	OngoingHttpServerRequests     *ebpf.Map `ebpf:"ongoing_http_server_requests"`
	OngoingKafkaRequests          *ebpf.Map `ebpf:"ongoing_kafka_requests"`
	OngoingProduceMessages        *ebpf.Map `ebpf:"ongoing_produce_messages"`
	OngoingProduceTopics          *ebpf.Map `ebpf:"ongoing_produce_topics"`
	OngoingRedisRequests          *ebpf.Map `ebpf:"ongoing_redis_requests"`
	OngoingServerConnections      *ebpf.Map `ebpf:"ongoing_server_connections"`
	OngoingSqlQueries             *ebpf.Map `ebpf:"ongoing_sql_queries"`
	OngoingStreams                *ebpf.Map `ebpf:"ongoing_streams"`
	OutgoingTraceMap              *ebpf.Map `ebpf:"outgoing_trace_map"`
	ProduceRequests               *ebpf.Map `ebpf:"produce_requests"`
	ProduceTraceparents           *ebpf.Map `ebpf:"produce_traceparents"`
	RedisWrites                   *ebpf.Map `ebpf:"redis_writes"`
	TraceMap                      *ebpf.Map `ebpf:"trace_map"`
}

func (m *bpf_tpMaps) Close() error {
	return _Bpf_tpClose(
		m.Events,
		m.FetchRequests,
		m.FramerInvocationMap,
		m.GoOffsetsMap,
		m.GoTraceMap,
		m.GolangMapbucketStorageMap,
		m.GrpcFramerInvocationMap,
		m.HeaderReqMap,
		m.Http2ReqMap,
		m.IncomingTraceMap,
		m.KafkaRequests,
		m.Newproc1,
		m.OngoingClientConnections,
		m.OngoingGoHttp,
		m.OngoingGoroutines,
		m.OngoingGrpcClientRequests,
		m.OngoingGrpcHeaderWrites,
		m.OngoingGrpcOperateHeaders,
		m.OngoingGrpcRequestStatus,
		m.OngoingGrpcServerRequests,
		m.OngoingGrpcTransports,
		m.OngoingHttpClientRequests,
		m.OngoingHttpClientRequestsData,
		m.OngoingHttpServerRequests,
		m.OngoingKafkaRequests,
		m.OngoingProduceMessages,
		m.OngoingProduceTopics,
		m.OngoingRedisRequests,
		m.OngoingServerConnections,
		m.OngoingSqlQueries,
		m.OngoingStreams,
		m.OutgoingTraceMap,
		m.ProduceRequests,
		m.ProduceTraceparents,
		m.RedisWrites,
		m.TraceMap,
	)
}

// bpf_tpPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpPrograms struct {
	UprobeClientConnClose                     *ebpf.Program `ebpf:"uprobe_ClientConn_Close"`
	UprobeClientConnInvoke                    *ebpf.Program `ebpf:"uprobe_ClientConn_Invoke"`
	UprobeClientConnInvokeReturn              *ebpf.Program `ebpf:"uprobe_ClientConn_Invoke_return"`
	UprobeClientConnNewStream                 *ebpf.Program `ebpf:"uprobe_ClientConn_NewStream"`
	UprobeClientConnNewStreamReturn           *ebpf.Program `ebpf:"uprobe_ClientConn_NewStream_return"`
	UprobeServeHTTP                           *ebpf.Program `ebpf:"uprobe_ServeHTTP"`
	UprobeServeHTTPReturns                    *ebpf.Program `ebpf:"uprobe_ServeHTTPReturns"`
	UprobeClientStreamRecvMsgReturn           *ebpf.Program `ebpf:"uprobe_clientStream_RecvMsg_return"`
	UprobeClientRoundTrip                     *ebpf.Program `ebpf:"uprobe_client_roundTrip"`
	UprobeConnServe                           *ebpf.Program `ebpf:"uprobe_connServe"`
	UprobeConnServeRet                        *ebpf.Program `ebpf:"uprobe_connServeRet"`
	UprobeExecDC                              *ebpf.Program `ebpf:"uprobe_execDC"`
	UprobeGrpcFramerWriteHeaders              *ebpf.Program `ebpf:"uprobe_grpcFramerWriteHeaders"`
	UprobeGrpcFramerWriteHeadersReturns       *ebpf.Program `ebpf:"uprobe_grpcFramerWriteHeaders_returns"`
	UprobeHttp2FramerWriteHeaders             *ebpf.Program `ebpf:"uprobe_http2FramerWriteHeaders"`
	UprobeHttp2FramerWriteHeadersReturns      *ebpf.Program `ebpf:"uprobe_http2FramerWriteHeaders_returns"`
	UprobeHttp2ResponseWriterStateWriteHeader *ebpf.Program `ebpf:"uprobe_http2ResponseWriterStateWriteHeader"`
	UprobeHttp2RoundTrip                      *ebpf.Program `ebpf:"uprobe_http2RoundTrip"`
	UprobeHttp2ServerOperateHeaders           *ebpf.Program `ebpf:"uprobe_http2Server_operateHeaders"`
	UprobeHttp2serverConnRunHandler           *ebpf.Program `ebpf:"uprobe_http2serverConn_runHandler"`
	UprobeNetFdRead                           *ebpf.Program `ebpf:"uprobe_netFdRead"`
	UprobePersistConnRoundTrip                *ebpf.Program `ebpf:"uprobe_persistConnRoundTrip"`
	UprobeProcGoexit1                         *ebpf.Program `ebpf:"uprobe_proc_goexit1"`
	UprobeProcNewproc1                        *ebpf.Program `ebpf:"uprobe_proc_newproc1"`
	UprobeProcNewproc1Ret                     *ebpf.Program `ebpf:"uprobe_proc_newproc1_ret"`
	UprobeProtocolRoundtrip                   *ebpf.Program `ebpf:"uprobe_protocol_roundtrip"`
	UprobeProtocolRoundtripRet                *ebpf.Program `ebpf:"uprobe_protocol_roundtrip_ret"`
	UprobeQueryDC                             *ebpf.Program `ebpf:"uprobe_queryDC"`
	UprobeQueryReturn                         *ebpf.Program `ebpf:"uprobe_queryReturn"`
	UprobeReadRequestReturns                  *ebpf.Program `ebpf:"uprobe_readRequestReturns"`
	UprobeReadRequestStart                    *ebpf.Program `ebpf:"uprobe_readRequestStart"`
	UprobeReaderRead                          *ebpf.Program `ebpf:"uprobe_reader_read"`
	UprobeReaderReadRet                       *ebpf.Program `ebpf:"uprobe_reader_read_ret"`
	UprobeReaderSendMessage                   *ebpf.Program `ebpf:"uprobe_reader_send_message"`
	UprobeRedisProcess                        *ebpf.Program `ebpf:"uprobe_redis_process"`
	UprobeRedisProcessRet                     *ebpf.Program `ebpf:"uprobe_redis_process_ret"`
	UprobeRedisWithWriter                     *ebpf.Program `ebpf:"uprobe_redis_with_writer"`
	UprobeRedisWithWriterRet                  *ebpf.Program `ebpf:"uprobe_redis_with_writer_ret"`
	UprobeRoundTrip                           *ebpf.Program `ebpf:"uprobe_roundTrip"`
	UprobeRoundTripReturn                     *ebpf.Program `ebpf:"uprobe_roundTripReturn"`
	UprobeSaramaBrokerWrite                   *ebpf.Program `ebpf:"uprobe_sarama_broker_write"`
	UprobeSaramaResponsePromiseHandle         *ebpf.Program `ebpf:"uprobe_sarama_response_promise_handle"`
	UprobeSaramaSendInternal                  *ebpf.Program `ebpf:"uprobe_sarama_sendInternal"`
	UprobeServerHandleStream                  *ebpf.Program `ebpf:"uprobe_server_handleStream"`
	UprobeServerHandleStreamReturn            *ebpf.Program `ebpf:"uprobe_server_handleStream_return"`
	UprobeServerHandlerTransportHandleStreams *ebpf.Program `ebpf:"uprobe_server_handler_transport_handle_streams"`
	UprobeTransportHttp2ClientNewStream       *ebpf.Program `ebpf:"uprobe_transport_http2Client_NewStream"`
	UprobeTransportWriteStatus                *ebpf.Program `ebpf:"uprobe_transport_writeStatus"`
	UprobeWriteSubset                         *ebpf.Program `ebpf:"uprobe_writeSubset"`
	UprobeWriterProduce                       *ebpf.Program `ebpf:"uprobe_writer_produce"`
	UprobeWriterWriteMessages                 *ebpf.Program `ebpf:"uprobe_writer_write_messages"`
}

func (p *bpf_tpPrograms) Close() error {
	return _Bpf_tpClose(
		p.UprobeClientConnClose,
		p.UprobeClientConnInvoke,
		p.UprobeClientConnInvokeReturn,
		p.UprobeClientConnNewStream,
		p.UprobeClientConnNewStreamReturn,
		p.UprobeServeHTTP,
		p.UprobeServeHTTPReturns,
		p.UprobeClientStreamRecvMsgReturn,
		p.UprobeClientRoundTrip,
		p.UprobeConnServe,
		p.UprobeConnServeRet,
		p.UprobeExecDC,
		p.UprobeGrpcFramerWriteHeaders,
		p.UprobeGrpcFramerWriteHeadersReturns,
		p.UprobeHttp2FramerWriteHeaders,
		p.UprobeHttp2FramerWriteHeadersReturns,
		p.UprobeHttp2ResponseWriterStateWriteHeader,
		p.UprobeHttp2RoundTrip,
		p.UprobeHttp2ServerOperateHeaders,
		p.UprobeHttp2serverConnRunHandler,
		p.UprobeNetFdRead,
		p.UprobePersistConnRoundTrip,
		p.UprobeProcGoexit1,
		p.UprobeProcNewproc1,
		p.UprobeProcNewproc1Ret,
		p.UprobeProtocolRoundtrip,
		p.UprobeProtocolRoundtripRet,
		p.UprobeQueryDC,
		p.UprobeQueryReturn,
		p.UprobeReadRequestReturns,
		p.UprobeReadRequestStart,
		p.UprobeReaderRead,
		p.UprobeReaderReadRet,
		p.UprobeReaderSendMessage,
		p.UprobeRedisProcess,
		p.UprobeRedisProcessRet,
		p.UprobeRedisWithWriter,
		p.UprobeRedisWithWriterRet,
		p.UprobeRoundTrip,
		p.UprobeRoundTripReturn,
		p.UprobeSaramaBrokerWrite,
		p.UprobeSaramaResponsePromiseHandle,
		p.UprobeSaramaSendInternal,
		p.UprobeServerHandleStream,
		p.UprobeServerHandleStreamReturn,
		p.UprobeServerHandlerTransportHandleStreams,
		p.UprobeTransportHttp2ClientNewStream,
		p.UprobeTransportWriteStatus,
		p.UprobeWriteSubset,
		p.UprobeWriterProduce,
		p.UprobeWriterWriteMessages,
	)
}

func _Bpf_tpClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_tp_arm64_bpfel.o
var _Bpf_tpBytes []byte
