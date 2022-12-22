package v20

import "context"

// 可以发起 RPC 调用的客户端
type I9RPCClient interface {
	// 发起 RPC 调用
	F8SendRPC(i9ctx context.Context, p7s6req *S6RPCRequest) (*S6RPCResponse, error)
}

// RPC 的请求
type S6RPCRequest struct {
	ServiceName             string
	FunctionName            string
	FunctionInputEncodeData []byte
}

// RPC 的响应
type S6RPCResponse struct {
	FunctionOutputEncodeData []byte
}
