package main

import (
	v20 "micro-service-go/v20"
	"micro-service-go/v20/protocol"
)

func main() {
	p7s6server := v20.F8NewS6RPCServer(
		v20.F8SetS6RPCServerProtocol(protocol.F8NewS6CustomRPC()),
	)
	p7s6UserService := &v20.S6UserService{}
	p7s6server.F8RegisterService(p7s6UserService)

	_ = p7s6server.F8Start("127.0.0.1:9602")
}
