// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package generictracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_tp_debugCallProtocolArgsT struct {
	PidConn    bpf_tp_debugPidConnectionInfoT
	SmallBuf   [24]uint8
	U_buf      uint64
	BytesLen   int32
	Ssl        uint8
	Direction  uint8
	OrigDport  uint16
	PacketType uint8
	_          [7]byte
}

type bpf_tp_debugConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpf_tp_debugCpSupportDataT struct {
	T_key      bpf_tp_debugTraceKeyT
	RealClient uint8
}

type bpf_tp_debugEgressKeyT struct {
	S_port uint16
	D_port uint16
}

type bpf_tp_debugGrpcFramesCtxT struct {
	PrevInfo        bpf_tp_debugHttp2GrpcRequestT
	HasPrevInfo     uint8
	_               [3]byte
	Pos             int32
	SavedBufPos     int32
	SavedStreamId   uint32
	FoundDataFrame  uint8
	Iterations      uint8
	TerminateSearch uint8
	_               [1]byte
	Stream          bpf_tp_debugHttp2ConnStreamT
	Args            bpf_tp_debugCallProtocolArgsT
}

type bpf_tp_debugHttp2ConnStreamT struct {
	PidConn  bpf_tp_debugPidConnectionInfoT
	StreamId uint32
}

type bpf_tp_debugHttp2GrpcRequestT struct {
	Flags           uint8
	_               [3]byte
	ConnInfo        bpf_tp_debugConnectionInfoT
	Data            [256]uint8
	RetData         [64]uint8
	Type            uint8
	_               [3]byte
	Len             int32
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Ssl     uint8
	NewConn uint8
	_       [2]byte
	Tp      struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
}

type bpf_tp_debugHttpConnectionMetadataT struct {
	Pid struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Type uint8
}

type bpf_tp_debugHttpInfoT struct {
	Flags           uint8
	_               [3]byte
	ConnInfo        bpf_tp_debugConnectionInfoT
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [256]uint8
	Len             uint32
	RespLen         uint32
	Status          uint16
	Type            uint8
	Ssl             uint8
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	ExtraId uint64
	TaskTid uint32
	_       [4]byte
}

type bpf_tp_debugMsgBufferT struct {
	Buf [256]uint8
	Pos uint16
}

type bpf_tp_debugPartialConnectionInfoT struct {
	S_addr [16]uint8
	S_port uint16
	D_port uint16
	TcpSeq uint32
}

type bpf_tp_debugPidConnectionInfoT struct {
	Conn bpf_tp_debugConnectionInfoT
	Pid  uint32
}

type bpf_tp_debugPidKeyT struct {
	Pid uint32
	Ns  uint32
}

type bpf_tp_debugRecvArgsT struct {
	SockPtr  uint64
	IovecCtx [40]uint8
}

type bpf_tp_debugSendArgsT struct {
	P_conn  bpf_tp_debugPidConnectionInfoT
	Size    uint64
	SockPtr uint64
}

type bpf_tp_debugSockArgsT struct {
	Addr       uint64
	AcceptTime uint64
}

type bpf_tp_debugSslArgsT struct {
	Ssl    uint64
	Buf    uint64
	LenPtr uint64
}

type bpf_tp_debugSslPidConnectionInfoT struct {
	P_conn    bpf_tp_debugPidConnectionInfoT
	OrigDport uint16
	_         [2]byte
}

type bpf_tp_debugTcpReqT struct {
	Flags           uint8
	_               [3]byte
	ConnInfo        bpf_tp_debugConnectionInfoT
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [256]uint8
	Rbuf            [128]uint8
	Len             uint32
	RespLen         uint32
	Ssl             uint8
	Direction       uint8
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	_  [2]byte
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	ExtraId uint64
}

type bpf_tp_debugTpInfoPidT struct {
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	Pid     uint32
	Valid   uint8
	ReqType uint8
	_       [2]byte
}

type bpf_tp_debugTraceKeyT struct {
	P_key   bpf_tp_debugPidKeyT
	ExtraId uint64
}

type bpf_tp_debugTraceMapKeyT struct {
	Conn bpf_tp_debugConnectionInfoT
	Type uint32
}

// loadBpf_tp_debug returns the embedded CollectionSpec for bpf_tp_debug.
func loadBpf_tp_debug() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_tp_debugBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_tp_debug: %w", err)
	}

	return spec, err
}

// loadBpf_tp_debugObjects loads bpf_tp_debug and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_tp_debugObjects
//	*bpf_tp_debugPrograms
//	*bpf_tp_debugMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_tp_debugObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_tp_debug()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_tp_debugSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugSpecs struct {
	bpf_tp_debugProgramSpecs
	bpf_tp_debugMapSpecs
}

// bpf_tp_debugSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugProgramSpecs struct {
	BeylaAsyncReset                        *ebpf.ProgramSpec `ebpf:"beyla_async_reset"`
	BeylaEmitAsyncInit                     *ebpf.ProgramSpec `ebpf:"beyla_emit_async_init"`
	BeylaKprobeSockRecvmsg                 *ebpf.ProgramSpec `ebpf:"beyla_kprobe_sock_recvmsg"`
	BeylaKprobeSysExit                     *ebpf.ProgramSpec `ebpf:"beyla_kprobe_sys_exit"`
	BeylaKprobeTcpCleanupRbuf              *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_cleanup_rbuf"`
	BeylaKprobeTcpClose                    *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_close"`
	BeylaKprobeTcpConnect                  *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_connect"`
	BeylaKprobeTcpRateCheckAppLimited      *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_rate_check_app_limited"`
	BeylaKprobeTcpRcvEstablished           *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_rcv_established"`
	BeylaKprobeTcpRecvmsg                  *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_recvmsg"`
	BeylaKprobeTcpSendmsg                  *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_sendmsg"`
	BeylaKprobeUnixStreamRecvmsg           *ebpf.ProgramSpec `ebpf:"beyla_kprobe_unix_stream_recvmsg"`
	BeylaKprobeUnixStreamSendmsg           *ebpf.ProgramSpec `ebpf:"beyla_kprobe_unix_stream_sendmsg"`
	BeylaKretprobeSockAlloc                *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_sock_alloc"`
	BeylaKretprobeSockRecvmsg              *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_sock_recvmsg"`
	BeylaKretprobeSysAccept4               *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_sys_accept4"`
	BeylaKretprobeSysClone                 *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_sys_clone"`
	BeylaKretprobeSysConnect               *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_sys_connect"`
	BeylaKretprobeTcpRecvmsg               *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_tcp_recvmsg"`
	BeylaKretprobeTcpSendmsg               *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_tcp_sendmsg"`
	BeylaKretprobeUnixStreamRecvmsg        *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_unix_stream_recvmsg"`
	BeylaKretprobeUnixStreamSendmsg        *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_unix_stream_sendmsg"`
	BeylaProtocolHttp                      *ebpf.ProgramSpec `ebpf:"beyla_protocol_http"`
	BeylaProtocolHttp2                     *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2"`
	BeylaProtocolHttp2GrpcFrames           *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2_grpc_frames"`
	BeylaProtocolHttp2GrpcHandleEndFrame   *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2_grpc_handle_end_frame"`
	BeylaProtocolHttp2GrpcHandleStartFrame *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2_grpc_handle_start_frame"`
	BeylaProtocolTcp                       *ebpf.ProgramSpec `ebpf:"beyla_protocol_tcp"`
	BeylaSocketHttpFilter                  *ebpf.ProgramSpec `ebpf:"beyla_socket__http_filter"`
	BeylaUprobeSslRead                     *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ssl_read"`
	BeylaUprobeSslReadEx                   *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ssl_read_ex"`
	BeylaUprobeSslShutdown                 *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ssl_shutdown"`
	BeylaUprobeSslWrite                    *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ssl_write"`
	BeylaUprobeSslWriteEx                  *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ssl_write_ex"`
	BeylaUretprobeSslRead                  *ebpf.ProgramSpec `ebpf:"beyla_uretprobe_ssl_read"`
	BeylaUretprobeSslReadEx                *ebpf.ProgramSpec `ebpf:"beyla_uretprobe_ssl_read_ex"`
	BeylaUretprobeSslWrite                 *ebpf.ProgramSpec `ebpf:"beyla_uretprobe_ssl_write"`
	BeylaUretprobeSslWriteEx               *ebpf.ProgramSpec `ebpf:"beyla_uretprobe_ssl_write_ex"`
}

// bpf_tp_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugMapSpecs struct {
	ActiveAcceptArgs        *ebpf.MapSpec `ebpf:"active_accept_args"`
	ActiveConnectArgs       *ebpf.MapSpec `ebpf:"active_connect_args"`
	ActiveNodejsIds         *ebpf.MapSpec `ebpf:"active_nodejs_ids"`
	ActiveRecvArgs          *ebpf.MapSpec `ebpf:"active_recv_args"`
	ActiveSendArgs          *ebpf.MapSpec `ebpf:"active_send_args"`
	ActiveSendSockArgs      *ebpf.MapSpec `ebpf:"active_send_sock_args"`
	ActiveSslConnections    *ebpf.MapSpec `ebpf:"active_ssl_connections"`
	ActiveSslReadArgs       *ebpf.MapSpec `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs      *ebpf.MapSpec `ebpf:"active_ssl_write_args"`
	ActiveUnixSocks         *ebpf.MapSpec `ebpf:"active_unix_socks"`
	AsyncResetArgs          *ebpf.MapSpec `ebpf:"async_reset_args"`
	CloneMap                *ebpf.MapSpec `ebpf:"clone_map"`
	ConnectionMetaMem       *ebpf.MapSpec `ebpf:"connection_meta_mem"`
	CpSupportConnectInfo    *ebpf.MapSpec `ebpf:"cp_support_connect_info"`
	DebugEvents             *ebpf.MapSpec `ebpf:"debug_events"`
	Events                  *ebpf.MapSpec `ebpf:"events"`
	GrpcFramesCtxMem        *ebpf.MapSpec `ebpf:"grpc_frames_ctx_mem"`
	Http2InfoMem            *ebpf.MapSpec `ebpf:"http2_info_mem"`
	HttpInfoMem             *ebpf.MapSpec `ebpf:"http_info_mem"`
	IncomingTraceMap        *ebpf.MapSpec `ebpf:"incoming_trace_map"`
	IovecMem                *ebpf.MapSpec `ebpf:"iovec_mem"`
	JumpTable               *ebpf.MapSpec `ebpf:"jump_table"`
	MsgBuffers              *ebpf.MapSpec `ebpf:"msg_buffers"`
	NodejsParentMap         *ebpf.MapSpec `ebpf:"nodejs_parent_map"`
	OngoingHttp             *ebpf.MapSpec `ebpf:"ongoing_http"`
	OngoingHttp2Connections *ebpf.MapSpec `ebpf:"ongoing_http2_connections"`
	OngoingHttp2Grpc        *ebpf.MapSpec `ebpf:"ongoing_http2_grpc"`
	OngoingHttpFallback     *ebpf.MapSpec `ebpf:"ongoing_http_fallback"`
	OngoingTcpReq           *ebpf.MapSpec `ebpf:"ongoing_tcp_req"`
	OutgoingTraceMap        *ebpf.MapSpec `ebpf:"outgoing_trace_map"`
	PidCache                *ebpf.MapSpec `ebpf:"pid_cache"`
	PidTidToConn            *ebpf.MapSpec `ebpf:"pid_tid_to_conn"`
	ProtocolArgsMem         *ebpf.MapSpec `ebpf:"protocol_args_mem"`
	ServerTraces            *ebpf.MapSpec `ebpf:"server_traces"`
	SslToConn               *ebpf.MapSpec `ebpf:"ssl_to_conn"`
	SslToPidTid             *ebpf.MapSpec `ebpf:"ssl_to_pid_tid"`
	TcpConnectionMap        *ebpf.MapSpec `ebpf:"tcp_connection_map"`
	TcpReqMem               *ebpf.MapSpec `ebpf:"tcp_req_mem"`
	TpCharBufMem            *ebpf.MapSpec `ebpf:"tp_char_buf_mem"`
	TpInfoMem               *ebpf.MapSpec `ebpf:"tp_info_mem"`
	TraceMap                *ebpf.MapSpec `ebpf:"trace_map"`
	ValidPids               *ebpf.MapSpec `ebpf:"valid_pids"`
}

// bpf_tp_debugObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugObjects struct {
	bpf_tp_debugPrograms
	bpf_tp_debugMaps
}

func (o *bpf_tp_debugObjects) Close() error {
	return _Bpf_tp_debugClose(
		&o.bpf_tp_debugPrograms,
		&o.bpf_tp_debugMaps,
	)
}

// bpf_tp_debugMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugMaps struct {
	ActiveAcceptArgs        *ebpf.Map `ebpf:"active_accept_args"`
	ActiveConnectArgs       *ebpf.Map `ebpf:"active_connect_args"`
	ActiveNodejsIds         *ebpf.Map `ebpf:"active_nodejs_ids"`
	ActiveRecvArgs          *ebpf.Map `ebpf:"active_recv_args"`
	ActiveSendArgs          *ebpf.Map `ebpf:"active_send_args"`
	ActiveSendSockArgs      *ebpf.Map `ebpf:"active_send_sock_args"`
	ActiveSslConnections    *ebpf.Map `ebpf:"active_ssl_connections"`
	ActiveSslReadArgs       *ebpf.Map `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs      *ebpf.Map `ebpf:"active_ssl_write_args"`
	ActiveUnixSocks         *ebpf.Map `ebpf:"active_unix_socks"`
	AsyncResetArgs          *ebpf.Map `ebpf:"async_reset_args"`
	CloneMap                *ebpf.Map `ebpf:"clone_map"`
	ConnectionMetaMem       *ebpf.Map `ebpf:"connection_meta_mem"`
	CpSupportConnectInfo    *ebpf.Map `ebpf:"cp_support_connect_info"`
	DebugEvents             *ebpf.Map `ebpf:"debug_events"`
	Events                  *ebpf.Map `ebpf:"events"`
	GrpcFramesCtxMem        *ebpf.Map `ebpf:"grpc_frames_ctx_mem"`
	Http2InfoMem            *ebpf.Map `ebpf:"http2_info_mem"`
	HttpInfoMem             *ebpf.Map `ebpf:"http_info_mem"`
	IncomingTraceMap        *ebpf.Map `ebpf:"incoming_trace_map"`
	IovecMem                *ebpf.Map `ebpf:"iovec_mem"`
	JumpTable               *ebpf.Map `ebpf:"jump_table"`
	MsgBuffers              *ebpf.Map `ebpf:"msg_buffers"`
	NodejsParentMap         *ebpf.Map `ebpf:"nodejs_parent_map"`
	OngoingHttp             *ebpf.Map `ebpf:"ongoing_http"`
	OngoingHttp2Connections *ebpf.Map `ebpf:"ongoing_http2_connections"`
	OngoingHttp2Grpc        *ebpf.Map `ebpf:"ongoing_http2_grpc"`
	OngoingHttpFallback     *ebpf.Map `ebpf:"ongoing_http_fallback"`
	OngoingTcpReq           *ebpf.Map `ebpf:"ongoing_tcp_req"`
	OutgoingTraceMap        *ebpf.Map `ebpf:"outgoing_trace_map"`
	PidCache                *ebpf.Map `ebpf:"pid_cache"`
	PidTidToConn            *ebpf.Map `ebpf:"pid_tid_to_conn"`
	ProtocolArgsMem         *ebpf.Map `ebpf:"protocol_args_mem"`
	ServerTraces            *ebpf.Map `ebpf:"server_traces"`
	SslToConn               *ebpf.Map `ebpf:"ssl_to_conn"`
	SslToPidTid             *ebpf.Map `ebpf:"ssl_to_pid_tid"`
	TcpConnectionMap        *ebpf.Map `ebpf:"tcp_connection_map"`
	TcpReqMem               *ebpf.Map `ebpf:"tcp_req_mem"`
	TpCharBufMem            *ebpf.Map `ebpf:"tp_char_buf_mem"`
	TpInfoMem               *ebpf.Map `ebpf:"tp_info_mem"`
	TraceMap                *ebpf.Map `ebpf:"trace_map"`
	ValidPids               *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpf_tp_debugMaps) Close() error {
	return _Bpf_tp_debugClose(
		m.ActiveAcceptArgs,
		m.ActiveConnectArgs,
		m.ActiveNodejsIds,
		m.ActiveRecvArgs,
		m.ActiveSendArgs,
		m.ActiveSendSockArgs,
		m.ActiveSslConnections,
		m.ActiveSslReadArgs,
		m.ActiveSslWriteArgs,
		m.ActiveUnixSocks,
		m.AsyncResetArgs,
		m.CloneMap,
		m.ConnectionMetaMem,
		m.CpSupportConnectInfo,
		m.DebugEvents,
		m.Events,
		m.GrpcFramesCtxMem,
		m.Http2InfoMem,
		m.HttpInfoMem,
		m.IncomingTraceMap,
		m.IovecMem,
		m.JumpTable,
		m.MsgBuffers,
		m.NodejsParentMap,
		m.OngoingHttp,
		m.OngoingHttp2Connections,
		m.OngoingHttp2Grpc,
		m.OngoingHttpFallback,
		m.OngoingTcpReq,
		m.OutgoingTraceMap,
		m.PidCache,
		m.PidTidToConn,
		m.ProtocolArgsMem,
		m.ServerTraces,
		m.SslToConn,
		m.SslToPidTid,
		m.TcpConnectionMap,
		m.TcpReqMem,
		m.TpCharBufMem,
		m.TpInfoMem,
		m.TraceMap,
		m.ValidPids,
	)
}

// bpf_tp_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugPrograms struct {
	BeylaAsyncReset                        *ebpf.Program `ebpf:"beyla_async_reset"`
	BeylaEmitAsyncInit                     *ebpf.Program `ebpf:"beyla_emit_async_init"`
	BeylaKprobeSockRecvmsg                 *ebpf.Program `ebpf:"beyla_kprobe_sock_recvmsg"`
	BeylaKprobeSysExit                     *ebpf.Program `ebpf:"beyla_kprobe_sys_exit"`
	BeylaKprobeTcpCleanupRbuf              *ebpf.Program `ebpf:"beyla_kprobe_tcp_cleanup_rbuf"`
	BeylaKprobeTcpClose                    *ebpf.Program `ebpf:"beyla_kprobe_tcp_close"`
	BeylaKprobeTcpConnect                  *ebpf.Program `ebpf:"beyla_kprobe_tcp_connect"`
	BeylaKprobeTcpRateCheckAppLimited      *ebpf.Program `ebpf:"beyla_kprobe_tcp_rate_check_app_limited"`
	BeylaKprobeTcpRcvEstablished           *ebpf.Program `ebpf:"beyla_kprobe_tcp_rcv_established"`
	BeylaKprobeTcpRecvmsg                  *ebpf.Program `ebpf:"beyla_kprobe_tcp_recvmsg"`
	BeylaKprobeTcpSendmsg                  *ebpf.Program `ebpf:"beyla_kprobe_tcp_sendmsg"`
	BeylaKprobeUnixStreamRecvmsg           *ebpf.Program `ebpf:"beyla_kprobe_unix_stream_recvmsg"`
	BeylaKprobeUnixStreamSendmsg           *ebpf.Program `ebpf:"beyla_kprobe_unix_stream_sendmsg"`
	BeylaKretprobeSockAlloc                *ebpf.Program `ebpf:"beyla_kretprobe_sock_alloc"`
	BeylaKretprobeSockRecvmsg              *ebpf.Program `ebpf:"beyla_kretprobe_sock_recvmsg"`
	BeylaKretprobeSysAccept4               *ebpf.Program `ebpf:"beyla_kretprobe_sys_accept4"`
	BeylaKretprobeSysClone                 *ebpf.Program `ebpf:"beyla_kretprobe_sys_clone"`
	BeylaKretprobeSysConnect               *ebpf.Program `ebpf:"beyla_kretprobe_sys_connect"`
	BeylaKretprobeTcpRecvmsg               *ebpf.Program `ebpf:"beyla_kretprobe_tcp_recvmsg"`
	BeylaKretprobeTcpSendmsg               *ebpf.Program `ebpf:"beyla_kretprobe_tcp_sendmsg"`
	BeylaKretprobeUnixStreamRecvmsg        *ebpf.Program `ebpf:"beyla_kretprobe_unix_stream_recvmsg"`
	BeylaKretprobeUnixStreamSendmsg        *ebpf.Program `ebpf:"beyla_kretprobe_unix_stream_sendmsg"`
	BeylaProtocolHttp                      *ebpf.Program `ebpf:"beyla_protocol_http"`
	BeylaProtocolHttp2                     *ebpf.Program `ebpf:"beyla_protocol_http2"`
	BeylaProtocolHttp2GrpcFrames           *ebpf.Program `ebpf:"beyla_protocol_http2_grpc_frames"`
	BeylaProtocolHttp2GrpcHandleEndFrame   *ebpf.Program `ebpf:"beyla_protocol_http2_grpc_handle_end_frame"`
	BeylaProtocolHttp2GrpcHandleStartFrame *ebpf.Program `ebpf:"beyla_protocol_http2_grpc_handle_start_frame"`
	BeylaProtocolTcp                       *ebpf.Program `ebpf:"beyla_protocol_tcp"`
	BeylaSocketHttpFilter                  *ebpf.Program `ebpf:"beyla_socket__http_filter"`
	BeylaUprobeSslRead                     *ebpf.Program `ebpf:"beyla_uprobe_ssl_read"`
	BeylaUprobeSslReadEx                   *ebpf.Program `ebpf:"beyla_uprobe_ssl_read_ex"`
	BeylaUprobeSslShutdown                 *ebpf.Program `ebpf:"beyla_uprobe_ssl_shutdown"`
	BeylaUprobeSslWrite                    *ebpf.Program `ebpf:"beyla_uprobe_ssl_write"`
	BeylaUprobeSslWriteEx                  *ebpf.Program `ebpf:"beyla_uprobe_ssl_write_ex"`
	BeylaUretprobeSslRead                  *ebpf.Program `ebpf:"beyla_uretprobe_ssl_read"`
	BeylaUretprobeSslReadEx                *ebpf.Program `ebpf:"beyla_uretprobe_ssl_read_ex"`
	BeylaUretprobeSslWrite                 *ebpf.Program `ebpf:"beyla_uretprobe_ssl_write"`
	BeylaUretprobeSslWriteEx               *ebpf.Program `ebpf:"beyla_uretprobe_ssl_write_ex"`
}

func (p *bpf_tp_debugPrograms) Close() error {
	return _Bpf_tp_debugClose(
		p.BeylaAsyncReset,
		p.BeylaEmitAsyncInit,
		p.BeylaKprobeSockRecvmsg,
		p.BeylaKprobeSysExit,
		p.BeylaKprobeTcpCleanupRbuf,
		p.BeylaKprobeTcpClose,
		p.BeylaKprobeTcpConnect,
		p.BeylaKprobeTcpRateCheckAppLimited,
		p.BeylaKprobeTcpRcvEstablished,
		p.BeylaKprobeTcpRecvmsg,
		p.BeylaKprobeTcpSendmsg,
		p.BeylaKprobeUnixStreamRecvmsg,
		p.BeylaKprobeUnixStreamSendmsg,
		p.BeylaKretprobeSockAlloc,
		p.BeylaKretprobeSockRecvmsg,
		p.BeylaKretprobeSysAccept4,
		p.BeylaKretprobeSysClone,
		p.BeylaKretprobeSysConnect,
		p.BeylaKretprobeTcpRecvmsg,
		p.BeylaKretprobeTcpSendmsg,
		p.BeylaKretprobeUnixStreamRecvmsg,
		p.BeylaKretprobeUnixStreamSendmsg,
		p.BeylaProtocolHttp,
		p.BeylaProtocolHttp2,
		p.BeylaProtocolHttp2GrpcFrames,
		p.BeylaProtocolHttp2GrpcHandleEndFrame,
		p.BeylaProtocolHttp2GrpcHandleStartFrame,
		p.BeylaProtocolTcp,
		p.BeylaSocketHttpFilter,
		p.BeylaUprobeSslRead,
		p.BeylaUprobeSslReadEx,
		p.BeylaUprobeSslShutdown,
		p.BeylaUprobeSslWrite,
		p.BeylaUprobeSslWriteEx,
		p.BeylaUretprobeSslRead,
		p.BeylaUretprobeSslReadEx,
		p.BeylaUretprobeSslWrite,
		p.BeylaUretprobeSslWriteEx,
	)
}

func _Bpf_tp_debugClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_tp_debug_arm64_bpfel.o
var _Bpf_tp_debugBytes []byte
