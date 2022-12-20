package v20

import "context"

type I9RPC interface {
	F8GetServiceName() string
}

type S6AnyRequest struct {
	Body string `json:"msg"`
}

type S6AnyResponse struct {
	Body string `json:"msg"`
}

type I9Proxy interface {
	F8Invoke(i9ctx context.Context, req *S6AnyRequest) (*S6AnyResponse, error)
}

func f8CoverWithRPC(i9rpc I9RPC) {

}
