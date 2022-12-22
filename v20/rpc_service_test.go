package v20

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type s6MockI9RPCClient struct {
	p7s6t    *testing.T
	p7s6req  *S6RPCRequest
	p7s6resp *S6RPCResponse
	err      error
}

func (p7this *s6MockI9RPCClient) F8SendRPC(i9ctx context.Context, req *S6RPCRequest) (*S6RPCResponse, error) {
	assert.Equal(p7this.p7s6t, p7this.p7s6req, req)
	return p7this.p7s6resp, p7this.err
}

type s6MockI9RPCService struct {
	i9RPCService I9RPCService
	f8SendRPC    func() (any, error)
}

func TestF8CoverWithRPC(p7s6t *testing.T) {
	s5s6case := []struct {
		name            string
		p7s6MockClient  *s6MockI9RPCClient
		p7s6MockService *s6MockI9RPCService
		wantResp        *S6F8GetUserByIdResponse
		wantErr         error
	}{
		{
			name: "user_rpc_service_client",
			p7s6MockClient: &s6MockI9RPCClient{
				p7s6t: p7s6t,
				p7s6req: &S6RPCRequest{
					ServiceName:             "user-rpc-service",
					FunctionName:            "F8GetUserById",
					FunctionInputEncodeData: []byte(`{"userId":11}`),
				},
				p7s6resp: &S6RPCResponse{
					FunctionOutputEncodeData: []byte(`{"userId":11,"userName":"aa"}`),
				},
			},
			p7s6MockService: func() *s6MockI9RPCService {
				p7RPCServiceClient := &S6UserRPCService{}
				return &s6MockI9RPCService{
					i9RPCService: p7RPCServiceClient,
					f8SendRPC: func() (any, error) {
						return p7RPCServiceClient.F8GetUserById(context.Background(), &S6F8GetUserByIdRequest{UserId: 11})
					},
				}
			}(),
			wantResp: &S6F8GetUserByIdResponse{
				UserId:   11,
				UserName: "aa",
			},
		},
	}

	for _, s6case := range s5s6case {
		p7s6t.Run(s6case.name, func(p7s6t2 *testing.T) {
			F8CoverWithRPC(s6case.p7s6MockService.i9RPCService, s6case.p7s6MockClient)
			resp, err := s6case.p7s6MockService.f8SendRPC()
			assert.Equal(p7s6t2, s6case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, s6case.wantResp, resp)
		})
	}
}
