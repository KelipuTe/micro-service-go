package v20

import "context"

type S6UserServiceClient struct {
	F8GetUserById func(i9ctx context.Context, p7req *S6AnyRequest) (*S6AnyResponse, error)
}

func (p7this *S6UserServiceClient) F8GetServiceName() string {
	return "user-service"
}
