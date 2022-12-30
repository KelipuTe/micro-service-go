package protocol

// RPC 的请求数据
type S6RPCRequest struct {
	ServiceName             string
	FunctionName            string
	FunctionInputEncodeData []byte
}

// RPC 的响应数据
type S6RPCResponse struct {
	FunctionOutputEncodeData []byte
}
