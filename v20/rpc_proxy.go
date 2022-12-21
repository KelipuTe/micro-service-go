package v20

import (
	"context"
	"encoding/json"
	"reflect"
)

// 实现这个接口，表示结构体可以被改造成 RPC 服务
// 可以被改造成 RPC 服务的结构体里面应该都是方法
type I9RPCService interface {
	F8GetServiceName() string
}

// 可以发起 RPC 调用的客户端
type I9RPCClient interface {
	// 发起 RPC 调用
	F8DoRPC(i9ctx context.Context, req *S6Request) (*S6Response, error)
}

// F8CoverWithRPC 把结构体改造成 RPC 服务
// i9RPCService 一个可以被改造成 RPC 服务的结构体
// i9client 可以发起 RPC 调用的客户端
func F8CoverWithRPC(i9RPCService I9RPCService, i9RPCClient I9RPCClient) {
	// 这里肯定是拿到一个接口（结构体指针）
	i9RPCServiceValue := reflect.ValueOf(i9RPCService)
	// 通过结构体指针拿到结构体值
	s6RPCServiceValue := i9RPCServiceValue.Elem()
	// 通过结构体值拿到结构体类型
	s6RPCServiceType := s6RPCServiceValue.Type()

	// 这里应该全部都是方法
	s6RPCServiceFieldNum := s6RPCServiceType.NumField()
	for i := 0; i < s6RPCServiceFieldNum; i++ {
		// 拿到结构体属性
		s6StructField := s6RPCServiceType.Field(i)
		// 拿到结构体属性的值
		s6StructFieldValue := s6RPCServiceValue.Field(i)
		// 拿到结构体属性的类型
		s6StructFieldType := s6StructField.Type
		// 判断一下结构体属性是否可修改
		if !s6StructFieldValue.CanSet() {
			continue
		}
		// 用原来的本地方法构造一个新的 RPC 调用的方法
		f8NewFunc := func(args []reflect.Value) (results []reflect.Value) {
			// 处理方法的入参，这里只管第二个参数，第一个是 context
			input := args[1].Interface()
			// 处理方法的返回值，这里只管第一个参数，第二个是 error
			output := reflect.New(s6StructFieldType.Out(0).Elem()).Interface()
			inputEncode, err := json.Marshal(input)
			if err != nil {
				return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
			}
			req := &S6Request{
				ServiceName: i9RPCService.F8GetServiceName(),
				MethodName:  s6StructField.Name,
				Data:        inputEncode,
			}
			// 向远端发起调用
			resp, err := i9RPCClient.F8DoRPC(args[0].Interface().(context.Context), req)
			if err != nil {
				return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
			}
			err = json.Unmarshal(resp.Data, output)
			if err != nil {
				return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
			}
			return []reflect.Value{reflect.ValueOf(output), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
		}
		// 把原来的本地方法替换掉
		s6StructFieldValue.Set(reflect.MakeFunc(s6StructFieldType, f8NewFunc))
	}
}
