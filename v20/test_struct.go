package v20

import "context"

// 两边都要

type S6F8GetUserByIdRequest struct {
	UserId int `json:"userId"`
}

type S6F8GetUserByIdResponse struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
}

// client

type S6UserRPCService struct {
	F8GetUserById func(i9ctx context.Context, p7req *S6F8GetUserByIdRequest) (*S6F8GetUserByIdResponse, error)
}

func (p7this *S6UserRPCService) F8GetServiceName() string {
	return "user-rpc-service"
}

// server

type S6UserService struct {
}

func (p7this *S6UserService) F8GetServiceName() string {
	return "user-rpc-service"
}

func (p7this *S6UserService) F8GetUserById(i9ctx context.Context, p7req *S6F8GetUserByIdRequest) (*S6F8GetUserByIdResponse, error) {
	return &S6F8GetUserByIdResponse{
		UserId:   22,
		UserName: "bb",
	}, nil
}
