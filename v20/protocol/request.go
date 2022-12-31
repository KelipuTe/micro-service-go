package protocol

// RPC 的请求数据
type S6RPCRequest struct {
	ServiceName  string
	FunctionName string
	// 元数据
	M3MetaData map[string]string
	// 方法入参的序列化方式
	SerializeCode uint8
	// 方法入参
	FunctionInputDataEncode []byte
}

// RPC 的响应数据
type S6RPCResponse struct {
	// 返回的异常
	Error error
	// 方法出参的序列化方式
	SerializeCode uint8
	// 方法出参
	FunctionOutputDataEncode []byte
}
