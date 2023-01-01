package v20

import (
	"context"
	"github.com/stretchr/testify/assert"
	"micro-service-go/v20/protocol"
	"testing"
)

func TestF8HandleRPC(p7s6t *testing.T) {
	s5s6case := []struct {
		name           string
		p7s6RPCServer  *S6RPCServer
		p7s6RPCRequest *protocol.S6RPCRequest
		wantResp       *protocol.S6RPCResponse
		wantErr        error
	}{
		{
			name: "user_rpc_service_client",
			p7s6RPCServer: func() *S6RPCServer {
				p7s6server := F8NewS6RPCServer()
				p7s6UserService := &S6UserService{}
				p7s6server.F8RegisterService(p7s6UserService)
				return p7s6server
			}(),
			p7s6RPCRequest: &protocol.S6RPCRequest{
				ServiceName:             "user-rpc-service",
				FunctionName:            "F8GetUserById",
				FunctionInputDataEncode: []byte(`{"userId":22}`),
			},
			wantResp: &protocol.S6RPCResponse{
				FunctionOutputDataEncode: []byte(`{"userId":22,"userName":"bb"}`),
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			resp := &protocol.S6RPCResponse{}
			functionOutputDataEncode, err := t4value.p7s6RPCServer.F8HandleRPC(context.Background(), t4value.p7s6RPCRequest)
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			resp.FunctionOutputDataEncode = functionOutputDataEncode
			assert.Equal(p7s6t2, t4value.wantResp, resp)
		})
	}
}
