package v20

import (
	"context"
	"micro-service-go/v20/protocol"
	"micro-service-go/v20/serialize"
	"net"
)

// #### type ####

// 可以发起 RPC 调用的客户端
type I9RPCClient interface {
	// 获取序列化
	F8GetI9Serialize() serialize.I9Serialize
	// 获取协议
	F8GetI9Protocol() protocol.I9Protocol
	// 发起 RPC 调用
	F8SendRPC(i9ctx context.Context, p7s6req *protocol.S6RPCRequest) (*protocol.S6RPCResponse, error)
}

type S6RPCClient struct {
	// 序列化
	i9Serialize serialize.I9Serialize
	// 协议
	i9Protocol protocol.I9Protocol
}

// Option 设计模式
type F8S6RPCClientOption func(*S6RPCClient)

// #### func ####

func F8NewS6RPCClient(s5Option ...F8S6RPCClientOption) *S6RPCClient {
	p7s6client := &S6RPCClient{}
	for _, t4value := range s5Option {
		t4value(p7s6client)
	}
	if nil == p7s6client.i9Serialize {
		p7s6client.i9Serialize = serialize.F8NewS6Json()
	}
	if nil == p7s6client.i9Protocol {
		p7s6client.i9Protocol = protocol.F8NewS6Json()
	}
	return p7s6client
}

func F8SetS6RPCClientSerialize(i9Serializer serialize.I9Serialize) F8S6RPCClientOption {
	return func(p7this *S6RPCClient) {
		p7this.i9Serialize = i9Serializer
	}
}

func F8SetS6RPCClientProtocol(i9Protocol protocol.I9Protocol) F8S6RPCClientOption {
	return func(p7this *S6RPCClient) {
		p7this.i9Protocol = i9Protocol
	}
}

// #### struct func ####

func (p7this *S6RPCClient) F8GetI9Serialize() serialize.I9Serialize {
	return p7this.i9Serialize
}

func (p7this *S6RPCClient) F8GetI9Protocol() protocol.I9Protocol {
	return p7this.i9Protocol
}

func (p7this *S6RPCClient) F8SendRPC(i9ctx context.Context, p7s6req *protocol.S6RPCRequest) (*protocol.S6RPCResponse, error) {
	i9conn, _ := net.Dial("tcp", "127.0.0.1:9602")

	i9protocol := p7this.F8GetI9Protocol()
	s5ReqMsg, err := i9protocol.F8EncodeReq(p7s6req)
	if nil != err {
		return nil, err
	}

	_, _ = i9conn.Write(s5ReqMsg)
	s5RespMsg, err := i9protocol.F8ReadRespMsgFromTCP(i9conn)
	if err != nil {
		return nil, err
	}

	p7s6resp, err := i9protocol.F8DecodeResp(s5RespMsg)
	if err != nil {
		return nil, err
	}
	return p7s6resp, nil
}
