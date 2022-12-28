package main

import v20 "micro-service-go/v20"

func main() {
	p7s6server := v20.F8NewS6RPCServer()
	p7s6UserService := &v20.S6UserService{}
	p7s6server.F8RegisterService(p7s6UserService)
	_ = p7s6server.F8Start("127.0.0.1:9602")
}
