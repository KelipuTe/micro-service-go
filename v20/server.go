package v20

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
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

// 处理 rpc
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

func (p7this *S6RPCServer) F8Start(address string) error {
	i9listener, err := net.Listen("tcp", address)
	if nil != err {
		return err
	}
	for {
		i9conn, err2 := i9listener.Accept()
		if nil != err2 {
			log.Printf("accept with err: %s", err2)
		}
		go p7this.f8HandleTCP(i9conn)
	}
}

// 消息长度的长度
const c5LenOfMsgLen int = 8

func (p7this *S6RPCServer) f8HandleTCP(i9conn net.Conn) {
	for {
		s5ReqMsg, err := p7this.readFromTcp(i9conn)
		if err != nil {
			return
		}
		p7s6req := &S6RPCRequest{}
		_ = json.Unmarshal(s5ReqMsg, p7s6req)
		log.Printf("%s,%+v", string(s5ReqMsg), p7s6req)
		p7s6resp, err := p7this.F8HandleRPC(context.Background(), p7s6req)
		s5RespMsg, _ := json.Marshal(p7s6resp)
		log.Printf("%+v,%s", p7s6resp, string(s5RespMsg))
		err = p7this.writeToTcp(i9conn, s5RespMsg)
		if err != nil {
			fmt.Printf("sending response failed: %v", err)
		}
	}
}

func (p7this *S6RPCServer) readFromTcp(i9conn net.Conn) (s5msg []byte, err error) {
	s5MsgLen := make([]byte, c5LenOfMsgLen)
	readLen, err := i9conn.Read(s5MsgLen)
	defer func() {
		if err2 := recover(); err2 != nil {
			log.Printf("tcp connection panic with : %v", err2)
			err = errors.New(fmt.Sprintf("tcp connection panic with : %v", err2))
		}
	}()
	if err != nil {
		return nil, err
	}
	if c5LenOfMsgLen != readLen {
		return nil, errors.New("could not read the length data")
	}
	msgLen := binary.BigEndian.Uint64(s5MsgLen)
	s5msg = make([]byte, msgLen)
	_, err = io.ReadFull(i9conn, s5msg)
	return s5msg, err
}

func (p7this *S6RPCServer) writeToTcp(i9conn net.Conn, s5msg []byte) error {
	encode := make([]byte, c5LenOfMsgLen+len(s5msg))
	binary.BigEndian.PutUint64(encode[:c5LenOfMsgLen], uint64(len(s5msg)))
	copy(encode[8:], s5msg)
	i9conn.Write(encode)
	return nil
}
