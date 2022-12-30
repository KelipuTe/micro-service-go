package main

import (
	"context"
	"fmt"
	v20 "micro-service-go/v20"
)

func main() {
	p7s6client := v20.F8NewS6RPCClient()
	p7s6RPCService := &v20.S6UserRPCService{}
	v20.F8CoverWithRPC(p7s6client, p7s6RPCService)
	resp, err := p7s6RPCService.F8GetUserById(context.Background(), &v20.S6F8GetUserByIdRequest{UserId: 11})
	fmt.Println(resp, err)
}
