package v20

import (
	"context"
	"fmt"
	"reflect"
)

// 可以处理 RPC 调用的服务端
type S6RPCServer struct {
	// 本地服务
	m3service map[string]*S6ReflectService
}

func F8NewS6RPCServer() *S6RPCServer {
	return &S6RPCServer{
		m3service: make(map[string]*S6ReflectService, 4),
	}
}

// 注册本地服务
func (p7this *S6RPCServer) F8RegisterService(i9RPCService I9RPCService) {
	// 这里用本地服务对应的 RPC 服务的服务名作为 key
	// 这样就可以通过 RPC 客户端发过来的 RPC 调用里的服务名，找到对应的本地服务
	p7this.m3service[i9RPCService.F8GetServiceName()] = &S6ReflectService{
		i9RPCService:               i9RPCService,
		s6i9RPCServiceReflectValue: reflect.ValueOf(i9RPCService),
	}
}

func (p7this *S6RPCServer) F8HandleRPC(i9ctx context.Context, p7s6req *S6RPCRequest) (*S6RPCResponse, error) {
	p7s6resp := &S6RPCResponse{}
	p7s6service, ok := p7this.m3service[p7s6req.ServiceName]
	if !ok {
		return p7s6resp, fmt.Errorf("server: 未找到服务, 服务名 %s", p7s6req.ServiceName)
	}
	respData, err := p7s6service.f8HandleRPC(i9ctx, p7s6req.FunctionName, p7s6req.FunctionInputEncodeData)
	if err != nil {
		return p7s6resp, err
	}
	p7s6resp.FunctionOutputEncodeData = respData
	return p7s6resp, nil
}
