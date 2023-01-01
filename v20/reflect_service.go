package v20

import (
	"context"
	"micro-service-go/v20/serialize"
	"reflect"
)

// 本地服务的反射
type S6ReflectService struct {
	i9RPCService               I9RPCService
	s6i9RPCServiceReflectValue reflect.Value
}

func (p7this *S6ReflectService) f8HandleRPC(i9ctx context.Context, functionName string, i9serialize serialize.I9Serialize, functionInputEncodeData []byte) ([]byte, error) {
	// 通过方法名，从结构体的反射中找到方法
	s6MethodValue := p7this.s6i9RPCServiceReflectValue.MethodByName(functionName)
	// 拿到方法的第二个入参的类型，第一个是 context
	inputType := s6MethodValue.Type().In(1)
	// 构造方法的第二个入参参
	inputValue := reflect.New(inputType.Elem())
	input := inputValue.Interface()
	// 把传过来的编码后的入参解码，然后放到构造的入参上去
	err := i9serialize.F8Decode(functionInputEncodeData, input)
	if err != nil {
		return nil, err
	}
	output := s6MethodValue.Call([]reflect.Value{reflect.ValueOf(i9ctx), inputValue})
	// 判断有没有 error
	if len(output) > 1 && !output[1].IsZero() {
		return nil, output[1].Interface().(error)
	}
	return i9serialize.F8Encode(output[0].Interface())
}
