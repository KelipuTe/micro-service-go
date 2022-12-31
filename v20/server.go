package v20

import (
	"context"
	"fmt"
	"log"
	"micro-service-go/v20/protocol"
	"micro-service-go/v20/serialize"
	"net"
	"reflect"
)

// #### type ####

// 可以处理 RPC 调用的服务端
type S6RPCServer struct {
	// 序列化
	i9Serialize serialize.I9Serialize
	// 协议
	i9Protocol protocol.I9Protocol
	// 本地服务
	m3service map[string]*S6ReflectService

	i9listener net.Listener
}

// Option 设计模式
type F8S6RPCServerOption func(*S6RPCServer)

// #### func ####

func F8NewS6RPCServer(s5Option ...F8S6RPCServerOption) *S6RPCServer {
	p7s6server := &S6RPCServer{
		m3service: make(map[string]*S6ReflectService, 4),
	}
	for _, t4value := range s5Option {
		t4value(p7s6server)
	}
	if nil == p7s6server.i9Serialize {
		p7s6server.i9Serialize = serialize.F8NewS6Json()
	}
	if nil == p7s6server.i9Protocol {
		p7s6server.i9Protocol = protocol.F8NewS6Json()
	}
	return p7s6server
}

func F8SetS6RPCServerSerialize(i9Serializer serialize.I9Serialize) F8S6RPCServerOption {
	return func(p7this *S6RPCServer) {
		p7this.i9Serialize = i9Serializer
	}
}

func F8SetS6RPCServerProtocol(i9Protocol protocol.I9Protocol) F8S6RPCServerOption {
	return func(p7this *S6RPCServer) {
		p7this.i9Protocol = i9Protocol
	}
}

// #### struct func ####

// 注册本地服务
func (p7this *S6RPCServer) F8RegisterService(i9RPCService I9RPCService) {
	// 这里用本地服务对应的 RPC 服务的服务名作为 key
	// 这样就可以通过 RPC 客户端发过来的 RPC 调用里的服务名，找到对应的本地服务
	p7this.m3service[i9RPCService.F8GetServiceName()] = &S6ReflectService{
		i9RPCService:               i9RPCService,
		s6i9RPCServiceReflectValue: reflect.ValueOf(i9RPCService),
	}
}

// 处理 rpc
func (p7this *S6RPCServer) F8HandleRPC(i9ctx context.Context, p7s6req *protocol.S6RPCRequest) (*protocol.S6RPCResponse, error) {
	p7s6service, ok := p7this.m3service[p7s6req.ServiceName]
	if !ok {
		return nil, fmt.Errorf("service [%s] not found", p7s6req.ServiceName)
	}
	respData, err := p7s6service.f8HandleRPC(i9ctx, p7s6req.FunctionName, p7s6req.FunctionInputDataEncode)
	if nil != err {
		return nil, err
	}
	p7s6resp := &protocol.S6RPCResponse{}
	p7s6resp.FunctionOutputDataEncode = respData
	return p7s6resp, nil
}

func (p7this *S6RPCServer) f8HandleTCP(i9conn net.Conn) {
	for {
		s5ReqMsg, err := p7this.i9Protocol.F8ReadReqMsgFromTCP(i9conn)
		if err != nil {
			// 一旦从 TCP 读取数据发生异常，这个链接最好就是断掉，字节流的异常处理太麻烦了
			log.Printf("f8HandleTCP F8ReadReqMsgFromTCP with: %s", err)
			return
		}
		p7s6req, err := p7this.i9Protocol.F8DecodeReq(s5ReqMsg)
		if err != nil {
			log.Printf("f8HandleTCP F8DecodeReq with: %s", err)
		}
		p7s6resp, err := p7this.F8HandleRPC(context.Background(), p7s6req)
		if err != nil {
			log.Printf("f8HandleTCP F8HandleRPC with: %s", err)
		}
		s5RespMsg, err := p7this.i9Protocol.F8EncodeResp(p7s6resp)
		if err != nil {
			log.Printf("f8HandleTCP F8EncodeResp with: %s", err)
		}
		_, err = i9conn.Write(s5RespMsg)
		if err != nil {
			log.Printf("f8HandleTCP Write with: %s", err)
		}
	}
}

func (p7this *S6RPCServer) F8Start(address string) error {
	i9listener, err := net.Listen("tcp", address)
	if nil != err {
		return err
	}
	p7this.i9listener = i9listener
	for {
		i9conn, err2 := i9listener.Accept()
		if nil != err2 {
			log.Printf("F8Start Accept with : %s", err2)
		}
		go p7this.f8HandleTCP(i9conn)
	}
}
