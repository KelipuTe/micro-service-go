package v20

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type s6MockI9RPCService struct {
	i9RPCService I9RPCService
	f8DoRPC      func() (any, error)
}

type s6MockI9RPCClient struct {
	p7s6t    *testing.T
	p7s6req  *S6Request
	p7s6resp *S6Response
	err      error
}

func (p7this *s6MockI9RPCClient) F8DoRPC(i9ctx context.Context, req *S6Request) (*S6Response, error) {
	assert.Equal(p7this.p7s6t, p7this.p7s6req, req)
	return p7this.p7s6resp, p7this.err
}

func TestF8CoverWithRPC(p7s6t *testing.T) {
	s5s6case := []struct {
		name            string
		p7s6MockService *s6MockI9RPCService
		p7s6MockClient  *s6MockI9RPCClient
		wantResp        *S6AnyResponse
		wantErr         error
	}{
		{
			name: "user_rpc_service_client",
			p7s6MockService: func() *s6MockI9RPCService {
				p7RPCServiceClient := &S6UserRPCServiceClient{}
				return &s6MockI9RPCService{
					i9RPCService: p7RPCServiceClient,
					f8DoRPC: func() (any, error) {
						return p7RPCServiceClient.F8GetUserById(context.Background(), &S6AnyRequest{Body: "F8GetUserById"})
					},
				}
			}(),
			p7s6MockClient: &s6MockI9RPCClient{
				p7s6t: p7s6t,
				p7s6req: &S6Request{
					ServiceName: "user-service",
					MethodName:  "GetById",
					Data:        []byte(`{"msg":"这是GetById"}`),
				},
				p7s6resp: &S6Response{
					Data: []byte(`{"msg":"这是GetById的响应"}`),
				},
			},
			wantResp: &S6AnyResponse{
				Body: "这是GetById的响应",
			},
		},
	}

	for _, s6case := range s5s6case {
		p7s6t.Run(s6case.name, func(p7s6t2 *testing.T) {
			F8CoverWithRPC(s6case.p7s6MockService.i9RPCService, s6case.p7s6MockClient)
			resp, err := s6case.p7s6MockService.f8DoRPC()
			assert.Equal(p7s6t2, s6case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, s6case.wantResp, resp)
		})
	}
}
