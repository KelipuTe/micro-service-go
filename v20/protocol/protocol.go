package protocol

import "net"

type I9Protocol interface {
	// 将 RPC 的请求数据编码成协议规定的报文
	F8EncodeReq(p7s6req *S6RPCRequest) ([]byte, error)
	// 将协议规定的报文解码成 RPC 的请求数据
	F8DecodeReq(s5ReqMsg []byte) (*S6RPCRequest, error)
	// 将 RPC 的响应数据编码成协议规定的报文
	F8EncodeResp(p7s6resp *S6RPCResponse) ([]byte, error)
	// 将协议规定的报文解码成 RPC 的响应数据
	F8DecodeResp(s5RespMsg []byte) (*S6RPCResponse, error)
	// 从 TCP 连接中读取一条请求报文
	F8ReadReqMsgFromTCP(i9conn net.Conn) ([]byte, error)
	// 从 TCP 连接中读取一条响应报文
	F8ReadRespMsgFromTCP(i9conn net.Conn) ([]byte, error)
}
