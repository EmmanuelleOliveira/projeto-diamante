package client

import (
	"context"
	"testing"

	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/client/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	googlegrpc "google.golang.org/grpc"
)

type mockGrpc struct {
	mock.Mock
}

func (m *mockGrpc) CreateClient(ctx context.Context, in *pb.ClientRequest, _ ...googlegrpc.CallOption) *pb.ClientResponse {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.ClientResponse)
}

func TestService_CreateNewClient(t *testing.T) {
	cases := []struct {
		name        string
		req         *pb.ClientRequest
		expectedReq interface{}
		want        *pb.ClientResponse
		mockFunc    func(client *mockGrpc, req interface{})
	}{
		{
			name: "success create new client",
			req: &pb.ClientRequest{
				Name:           "Teste",
				Email:          "manu@gmail.com",
				DocumentNumber: "10025815411",
				PhoneNumber:    "87999991111",
				Cep:            "54406100",
			},
			expectedReq: &pb.ClientResponse{
				Client: &pb.Client{
					Id:             1,
					Name:           "Teste",
					Email:          "manu@gmail.com",
					DocumentNumber: "10025815411",
					PhoneNumber:    "87999991111",
					Cep:            "54406100",
					Address: &pb.Address{
						Street: "street test",
						City:   "city test",
						Uf:     "Test",
					},
				},
			},
			mockFunc: func(client *mockGrpc, req interface{}) {
				client.On("CreateCliente", mock.Anything, req).
					Return(&pb.ClientResponse{
						Client: &pb.Client{
							Id:             1,
							Name:           "Teste",
							Email:          "manu@gmail.com",
							DocumentNumber: "10025815411",
							PhoneNumber:    "87999991111",
							Cep:            "54406100",
							Address: &pb.Address{
								Street: "street test",
								City:   "city test",
								Uf:     "Test",
							},
						},
					}).Once()
			},
			want: &pb.ClientResponse{
				Client: &pb.Client{
					Id:             1,
					Name:           "Teste",
					Email:          "manu@gmail.com",
					DocumentNumber: "10025815411",
					PhoneNumber:    "87999991111",
					Cep:            "54406100",
					Address: &pb.Address{
						Street: "street test",
						City:   "city test",
						Uf:     "Test",
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := new(mockGrpc)
			result := client.CreateClient(context.Background(), tc.req)
			assert.Equal(t, tc.want, result)
		})
	}
}
