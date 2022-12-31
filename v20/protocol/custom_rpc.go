package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	// 自定义 RPC 协议头部长度的长度
	c5LenOfCustomRPCHeaderLen int = 4
	// 自定义 RPC 协议主体长度的长度
	c5LenOfCustomRPCBodyLen int = 4

	c5Version uint8 = 2
	c5ASCII10 byte  = '\n'
	c5ASCII13 byte  = '\r'
)

type S6CustomRPC struct {
	Version  uint8
	Compress uint8
}

func F8NewS6CustomRPC() S6CustomRPC {
	return S6CustomRPC{
		Version: c5Version,
	}
}

func (this S6CustomRPC) F8EncodeReq(p7s6req *S6RPCRequest) ([]byte, error) {
	headerLen := 11
	headerLen += len(p7s6req.ServiceName) + 2
	headerLen += len(p7s6req.FunctionName) + 2
	for t4key, t4value := range p7s6req.M3MetaData {
		headerLen += len(t4key) + len(t4value) + 3
	}
	bodyLen := len(p7s6req.FunctionInputDataEncode)

	s5MsgBody := make([]byte, headerLen+bodyLen)
	p7current := s5MsgBody

	// 请求头长度，4 个字节
	binary.BigEndian.PutUint32(p7current[:c5LenOfCustomRPCHeaderLen], uint32(headerLen))
	p7current = p7current[c5LenOfCustomRPCHeaderLen:]
	// 请求体长度，4 个字节
	binary.BigEndian.PutUint32(p7current[:c5LenOfCustomRPCBodyLen], uint32(bodyLen))
	p7current = p7current[c5LenOfCustomRPCBodyLen:]
	// 版本号，1 个字节
	p7current[0] = this.Version
	p7current = p7current[1:]
	// 序列化算法，1 个字节
	p7current[0] = p7s6req.SerializeCode
	p7current = p7current[1:]
	// 压缩算法，1 个字节
	p7current[0] = 0
	p7current = p7current[1:]

	// 请求数据
	copy(p7current, p7s6req.ServiceName)
	p7current = p7current[len(p7s6req.ServiceName):]
	p7current[0] = c5ASCII13
	p7current[1] = c5ASCII10
	p7current = p7current[2:]

	copy(p7current, p7s6req.FunctionName)
	p7current = p7current[len(p7s6req.FunctionName):]
	p7current[0] = c5ASCII13
	p7current[1] = c5ASCII10
	p7current = p7current[2:]

	for t4key, t4value := range p7s6req.M3MetaData {
		copy(p7current, t4key)
		p7current = p7current[len(t4key):]
		p7current[0] = ':'
		p7current = p7current[1:]
		copy(p7current, t4value)
		p7current = p7current[len(t4value):]
		p7current[0] = c5ASCII13
		p7current[1] = c5ASCII10
		p7current = p7current[2:]
	}

	copy(p7current, p7s6req.FunctionInputDataEncode)

	return s5MsgBody, nil
}

func (this S6CustomRPC) F8DecodeReq(s5ReqMsg []byte) (*S6RPCRequest, error) {
	p7s6req := &S6RPCRequest{}

	currentIndex := 0
	headerLen := binary.BigEndian.Uint32(s5ReqMsg[currentIndex : currentIndex+c5LenOfCustomRPCHeaderLen])
	currentIndex += c5LenOfCustomRPCHeaderLen
	//bodyLen := binary.BigEndian.Uint32(s5ReqMsg[currentIndex : currentIndex+c5LenOfCustomRPCBodyLen])
	currentIndex += c5LenOfCustomRPCBodyLen
	//version := s5ReqMsg[currentIndex+1]
	p7s6req.SerializeCode = s5ReqMsg[currentIndex+2]
	//CompressCode := s5ReqMsg[currentIndex+3]
	currentIndex += 3

	s5HeaderPart := s5ReqMsg[currentIndex:headerLen]
	t4index := bytes.Index(s5HeaderPart, []byte{c5ASCII13, c5ASCII10})
	p7s6req.ServiceName = string(s5HeaderPart[:t4index])
	currentIndex = t4index + 2

	s5HeaderPart = s5HeaderPart[currentIndex:]
	t4index = bytes.Index(s5HeaderPart, []byte{c5ASCII13, c5ASCII10})
	p7s6req.FunctionName = string(s5HeaderPart[:t4index])
	currentIndex = t4index + 2

	p7s6req.M3MetaData = make(map[string]string, 2)
	for {
		s5HeaderPart = s5HeaderPart[currentIndex:]
		t4index = bytes.Index(s5HeaderPart, []byte{c5ASCII13, c5ASCII10})
		if -1 == t4index {
			break
		}
		t4index2 := bytes.IndexByte(s5HeaderPart, ':')
		t4key := string(s5HeaderPart[:t4index2])
		t4value := string(s5HeaderPart[t4index2+1 : t4index])
		p7s6req.M3MetaData[t4key] = t4value
		currentIndex = t4index + 2
	}

	p7s6req.FunctionInputDataEncode = s5ReqMsg[headerLen:]

	return p7s6req, nil
}

func (this S6CustomRPC) F8EncodeResp(p7s6resp *S6RPCResponse) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (this S6CustomRPC) F8DecodeResp(s5RespMsg []byte) (*S6RPCResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (this S6CustomRPC) F8ReadReqMsgFromTCP(i9conn net.Conn) (s5ReqMsg []byte, err error) {
	defer func() {
		if err2 := recover(); nil != err2 {
			err = errors.New(fmt.Sprintf("tcp connection panic with : %v", err2))
		}
	}()

	s5ReqHeaderMsgLen := make([]byte, c5LenOfCustomRPCHeaderLen)
	readByteNum, err := i9conn.Read(s5ReqHeaderMsgLen)
	if nil != err {
		return nil, err
	}
	if c5LenOfCustomRPCHeaderLen != readByteNum {
		return nil, errors.New("could not read msg length")
	}
	reqHeaderMsgLen := binary.BigEndian.Uint32(s5ReqHeaderMsgLen)

	s5ReqBodyMsgLen := make([]byte, c5LenOfCustomRPCBodyLen)
	readByteNum, err = i9conn.Read(s5ReqBodyMsgLen)
	if nil != err {
		return nil, err
	}
	if c5LenOfCustomRPCBodyLen != readByteNum {
		return nil, errors.New("could not read msg length")
	}
	reqBodyMsgLen := binary.BigEndian.Uint32(s5ReqBodyMsgLen)

	s5ReqMsg = make([]byte, reqHeaderMsgLen+reqBodyMsgLen)
	// 这里需要把刚才读出来的 8 个字节放进去
	copy(s5ReqMsg, s5ReqHeaderMsgLen)
	copy(s5ReqMsg[4:], s5ReqBodyMsgLen)
	_, err = io.ReadFull(i9conn, s5ReqMsg[8:])
	return s5ReqMsg, err
}

func (this S6CustomRPC) F8ReadRespMsgFromTCP(i9conn net.Conn) (s5ReqMsg []byte, err error) {
	defer func() {
		if err2 := recover(); nil != err2 {
			err = errors.New(fmt.Sprintf("tcp connection panic with : %v", err2))
		}
	}()

	s5ReqHeaderMsgLen := make([]byte, c5LenOfCustomRPCHeaderLen)
	readByteNum, err := i9conn.Read(s5ReqHeaderMsgLen)
	if nil != err {
		return nil, err
	}
	if c5LenOfCustomRPCHeaderLen != readByteNum {
		return nil, errors.New("could not read msg length")
	}
	reqHeaderMsgLen := binary.BigEndian.Uint32(s5ReqHeaderMsgLen)

	s5ReqBodyMsgLen := make([]byte, c5LenOfCustomRPCBodyLen)
	readByteNum, err = i9conn.Read(s5ReqBodyMsgLen)
	if nil != err {
		return nil, err
	}
	if c5LenOfCustomRPCBodyLen != readByteNum {
		return nil, errors.New("could not read msg length")
	}
	reqBodyMsgLen := binary.BigEndian.Uint32(s5ReqBodyMsgLen)

	s5ReqMsg = make([]byte, reqHeaderMsgLen+reqBodyMsgLen)
	// 这里需要把刚才读出来的 8 个字节放进去
	copy(s5ReqMsg, s5ReqHeaderMsgLen)
	copy(s5ReqMsg[4:], s5ReqBodyMsgLen)
	_, err = io.ReadFull(i9conn, s5ReqMsg[8:])
	return s5ReqMsg, err
}
