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
)

// 可以发起 RPC 调用的客户端
type I9RPCClient interface {
	// 发起 RPC 调用
	F8SendRPC(i9ctx context.Context, p7s6req *S6RPCRequest) (*S6RPCResponse, error)
}

type S6RPCClient struct {
}

func (p7this *S6RPCClient) F8SendRPC(i9ctx context.Context, p7s6req *S6RPCRequest) (*S6RPCResponse, error) {
	i9conn, _ := net.Dial("tcp", "127.0.0.1:9602")
	s5msg, _ := json.Marshal(p7s6req)
	encode := make([]byte, c5LenOfMsgLen+len(s5msg))
	binary.BigEndian.PutUint64(encode[:c5LenOfMsgLen], uint64(len(s5msg)))
	copy(encode[8:], s5msg)
	i9conn.Write(encode)

	s5RespMsg, err := p7this.readFromTcp(i9conn)
	if err != nil {
		return nil, err
	}
	p7s6resp := &S6RPCResponse{}
	_ = json.Unmarshal(s5RespMsg, p7s6resp)
	return p7s6resp, nil
}

func (p7this *S6RPCClient) readFromTcp(i9conn net.Conn) (s5msg []byte, err error) {
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

// RPC 的请求
type S6RPCRequest struct {
	ServiceName             string
	FunctionName            string
	FunctionInputEncodeData []byte
}

// RPC 的响应
type S6RPCResponse struct {
	FunctionOutputEncodeData []byte
}
