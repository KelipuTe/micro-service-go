package v20

// RPC 的请求
type S6Request struct {
	ServiceName string
	MethodName  string
	Data        []byte
}

// RPC 的响应
type S6Response struct {
	Data []byte
}
