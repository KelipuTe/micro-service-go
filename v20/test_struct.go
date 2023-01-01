package v20

import (
	"context"
	"fmt"
	"log"
)

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

func (p7this *S6UserService) F8GetUserById(i9ctx context.Context, p7s6req *S6F8GetUserByIdRequest) (*S6F8GetUserByIdResponse, error) {
	log.Printf("flowId: %s", i9ctx.Value("flowId").(string))
	if 11 == p7s6req.UserId {
		return &S6F8GetUserByIdResponse{
			UserId:   11,
			UserName: "aa",
		}, nil
	} else if 22 == p7s6req.UserId {
		return &S6F8GetUserByIdResponse{
			UserId:   22,
			UserName: "bb",
		}, nil
	} else {
		return nil, fmt.Errorf("user id [%d] not found", p7s6req.UserId)
	}
}
