package v20

import (
	"context"
	"micro-service-go/v20/protocol"
	"reflect"
)

// 实现这个接口，表示本地服务可以被改造成 RPC 服务
// 可以被改造成 RPC 服务的本地服务（结构体）里面应该都是方法
type I9RPCService interface {
	// 获取本地服务对应的 RPC 服务的服务名
	F8GetServiceName() string
}

// F8CoverWithRPC 把结构体改造成 RPC 服务
// i9RPCClient 可以发起 RPC 调用的客户端
// i9RPCService 一个可以被改造成 RPC 服务的结构体
func F8CoverWithRPC(i9RPCClient I9RPCClient, i9RPCService I9RPCService) {
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
		// 用结构体原来的方法的信息构造一个新的 RPC 调用的方法
		f8NewFunc := func(args []reflect.Value) (results []reflect.Value) {
			// 处理方法的入参，这里只管第二个参数，第一个是 context
			input := args[1].Interface()
			// 处理方法的返回值，这里只管第一个参数，第二个是 error
			output := reflect.New(s6StructFieldType.Out(0).Elem()).Interface()
			// 把方法的入参序列化
			i9serialize := i9RPCClient.F8GetI9Serialize()
			inputEncode, err := i9serialize.F8Encode(input)
			if err != nil {
				return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
			}
			// 从 context 获取元数据
			m3ExtraData := map[string]string{}
			i9ctx := args[0].Interface()
			if i9ctxValue, ok := i9ctx.(context.Context); ok {
				m3ExtraData["flowId"] = i9ctxValue.Value("flowId").(string)
			}
			// 组装调用的请求数据
			p7s6req := &protocol.S6RPCRequest{
				ServiceName:             i9RPCService.F8GetServiceName(),
				FunctionName:            s6StructField.Name,
				M3MetaData:              m3ExtraData,
				SerializeCode:           i9serialize.F8GetCode(),
				FunctionInputDataEncode: inputEncode,
			}
			// 向远端发起调用
			resp, err := i9RPCClient.F8SendRPC(args[0].Interface().(context.Context), p7s6req)
			if err != nil {
				return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
			}
			if resp.Error != nil {
				return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(resp.Error)}
			}
			// 把返回的数据反序列化
			err = i9serialize.F8Decode(resp.FunctionOutputDataEncode, output)
			if err != nil {
				return []reflect.Value{reflect.ValueOf(output), reflect.ValueOf(err)}
			}
			return []reflect.Value{reflect.ValueOf(output), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
		}
		// 把结构体原来的方法替换成新构造的这个 RPC 调用的方法
		s6StructFieldValue.Set(reflect.MakeFunc(s6StructFieldType, f8NewFunc))
	}
}
