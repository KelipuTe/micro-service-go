package v20

import "context"

type S6AnyRequest struct {
	Body string `json:"msg"`
}

type S6AnyResponse struct {
	Body string `json:"msg"`
}

type S6UserRPCServiceClient struct {
	F8GetUserById func(i9ctx context.Context, p7req *S6AnyRequest) (*S6AnyResponse, error)
}

func (p7this *S6UserRPCServiceClient) F8GetServiceName() string {
	return "user-service"
}
