package main

import (
	"context"
	"fmt"
	"micro-service-go/v20"
	"micro-service-go/v20/protocol"
)

func main() {
	p7s6client := v20.F8NewS6RPCClient(
		v20.F8SetS6RPCClientProtocol(protocol.F8NewS6CustomRPC()),
	)
	p7s6RPCService := &v20.S6UserRPCService{}
	v20.F8CoverWithRPC(p7s6client, p7s6RPCService)

	i9ctx := context.Background()
	i9ctx = context.WithValue(i9ctx, "flowId", "flowId12345678")
	resp, err := p7s6RPCService.F8GetUserById(i9ctx, &v20.S6F8GetUserByIdRequest{UserId: 33})
	fmt.Println(resp, err)
}
