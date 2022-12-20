package v20

import (
	"testing"
)

func Test_f8CoverWithRPC(p7s6t *testing.T) {
	//s5s6case := []struct {
	//	name     string
	//	s        *s6MockService
	//	proxy    *s6MockProxy
	//	wantResp any
	//	wantErr  error
	//}{
	//	{
	//		name: "user service",
	//		s: func() *mockService {
	//			s := &UserServiceClient{}
	//			return &mockService{
	//				s: s,
	//				do: func() (any, error) {
	//					return s.GetById(context.Background(), &AnyRequest{Msg: "这是GetById"})
	//				},
	//			}
	//		}(),
	//		proxy: &mockProxy{
	//			t: t,
	//			req: &Request{
	//				ServiceName: "user-service",
	//				Method:      "GetById",
	//				Data:        []byte(`{"msg":"这是GetById"}`),
	//			},
	//			resp: &Response{
	//				Data: []byte(`{"msg":"这是GetById的响应"}`),
	//			},
	//		},
	//		wantResp: &AnyResponse{
	//			Msg: "这是GetById的响应",
	//		},
	//	},
	//}
	//
	//// 通过传入 mockProxy，然后我们执行一下方法上面的调用，
	//// 确保已经篡改成功了
	//for _, tc := range testCases {
	//	t.Run(tc.name, func(t *testing.T) {
	//		setFuncField(tc.s.s, tc.proxy)
	//		resp, err := tc.s.do()
	//		assert.Equal(t, tc.wantErr, err)
	//		if err != nil {
	//			return
	//		}
	//		assert.Equal(t, tc.wantResp, resp)
	//	})
	//}
}
