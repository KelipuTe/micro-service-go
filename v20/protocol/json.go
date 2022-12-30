package protocol

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

// 消息长度的长度
const c5LenOfJsonMsgLen int = 8

type S6Json struct {
}

func F8NewS6Json() *S6Json { return &S6Json{} }

func (this S6Json) F8EncodeReq(p7s6req *S6RPCRequest) ([]byte, error) {
	s5MsgBody, err := json.Marshal(p7s6req)
	if nil != err {
		return nil, err
	}
	log.Printf("F8EncodeReq:%s,%+v", string(s5MsgBody), p7s6req)
	s5ReqMsg := make([]byte, c5LenOfJsonMsgLen+len(s5MsgBody))
	binary.BigEndian.PutUint64(s5ReqMsg[:c5LenOfJsonMsgLen], uint64(len(s5MsgBody)))
	copy(s5ReqMsg[8:], s5MsgBody)
	return s5ReqMsg, nil
}

func (this S6Json) F8DecodeReq(s5ReqMsg []byte) (*S6RPCRequest, error) {
	p7s6req := &S6RPCRequest{}
	err := json.Unmarshal(s5ReqMsg, p7s6req)
	if nil != err {
		return nil, err
	}
	log.Printf("F8DecodeReq:%s,%+v", string(s5ReqMsg), p7s6req)
	return p7s6req, nil
}

func (this S6Json) F8EncodeResp(p7s6resp *S6RPCResponse) ([]byte, error) {
	s5MsgBody, err := json.Marshal(p7s6resp)
	if nil != err {
		return nil, err
	}
	log.Printf("F8EncodeReq:%s,%+v", string(s5MsgBody), p7s6resp)
	s5RespMsg := make([]byte, c5LenOfJsonMsgLen+len(s5MsgBody))
	binary.BigEndian.PutUint64(s5RespMsg[:c5LenOfJsonMsgLen], uint64(len(s5MsgBody)))
	copy(s5RespMsg[8:], s5MsgBody)
	return s5RespMsg, nil
}

func (this S6Json) F8DecodeResp(s5RespMsg []byte) (*S6RPCResponse, error) {
	p7s6resp := &S6RPCResponse{}
	err := json.Unmarshal(s5RespMsg, p7s6resp)
	if nil != err {
		return nil, err
	}
	log.Printf("F8DecodeResp:%s,%+v", string(s5RespMsg), p7s6resp)
	return p7s6resp, nil
}

func (this S6Json) F8ReadReqMsgFromTCP(i9conn net.Conn) (s5ReqMsg []byte, err error) {
	s5ReqMsgLen := make([]byte, c5LenOfJsonMsgLen)
	readByteNum, err := i9conn.Read(s5ReqMsgLen)
	defer func() {
		if err2 := recover(); nil != err2 {
			err = errors.New(fmt.Sprintf("tcp connection panic with : %v", err2))
		}
	}()
	if nil != err {
		return nil, err
	}
	if c5LenOfJsonMsgLen != readByteNum {
		return nil, errors.New("could not read msg length")
	}
	reqMsgLen := binary.BigEndian.Uint64(s5ReqMsgLen)
	s5ReqMsg = make([]byte, reqMsgLen)
	_, err = io.ReadFull(i9conn, s5ReqMsg)
	return s5ReqMsg, err
}

func (this S6Json) F8ReadRespMsgFromTCP(i9conn net.Conn) (s5RespMsg []byte, err error) {
	s5RespMsgLen := make([]byte, c5LenOfJsonMsgLen)
	readByteNum, err := i9conn.Read(s5RespMsgLen)
	defer func() {
		if err2 := recover(); nil != err2 {
			// 因为这个地方要返回异常，所以返回值要用命名的
			err = errors.New(fmt.Sprintf("tcp connection panic with : %v", err2))
		}
	}()
	if nil != err {
		return nil, err
	}
	if c5LenOfJsonMsgLen != readByteNum {
		return nil, errors.New("could not read msg length")
	}
	respMsgLen := binary.BigEndian.Uint64(s5RespMsgLen)
	s5RespMsg = make([]byte, respMsgLen)
	_, err = io.ReadFull(i9conn, s5RespMsg)
	return s5RespMsg, err
}
