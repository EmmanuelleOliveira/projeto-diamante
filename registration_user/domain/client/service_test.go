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

func TestService_CreateClient(t *testing.T) {
	cases := []struct {
		name        string
		req         *pb.ClientRequest
		expectedRes interface{}
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
			expectedRes: &pb.ClientResponse{
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
				Error: "",
			},
			mockFunc: func(client *mockGrpc, req interface{}) {
				client.On("CreateClient", mock.Anything, req).
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
						Error: "",
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
				Error: "",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := new(mockGrpc)
			tc.mockFunc(client, tc.req) //Configura o mock
			result := client.CreateClient(context.Background(), tc.req)
			assert.Equal(t, tc.want, result)
		})
	}
}

func (m *mockGrpc) GetAll(ctx context.Context, _ ...googlegrpc.CallOption) *pb.GetAllClientsResponse {
	args := m.Called(ctx)
	return args.Get(0).(*pb.GetAllClientsResponse)
}

func TestService_GetAll(t *testing.T) {
	cases := []struct {
		name     string
		want     *pb.GetAllClientsResponse
		mockFunc func(client *mockGrpc)
	}{
		{
			name: "success to get all clients",
			mockFunc: func(client *mockGrpc) {
				client.On("GetAll", mock.Anything).
					Return(&pb.GetAllClientsResponse{
						Clients: []*pb.Client{
							{
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
							{
								Id:             2,
								Name:           "Teste",
								Email:          "teste@gmail.com",
								DocumentNumber: "10025815412",
								PhoneNumber:    "81999991111",
								Cep:            "54406200",
								Address: &pb.Address{
									Street: "street test",
									City:   "city test",
									Uf:     "Test",
								},
							},
						},
						Error: "",
					}).Once()
			},
			want: &pb.GetAllClientsResponse{
				Clients: []*pb.Client{
					{
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
					{
						Id:             2,
						Name:           "Teste",
						Email:          "teste@gmail.com",
						DocumentNumber: "10025815412",
						PhoneNumber:    "81999991111",
						Cep:            "54406200",
						Address: &pb.Address{
							Street: "street test",
							City:   "city test",
							Uf:     "Test",
						},
					},
				},
				Error: "",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := new(mockGrpc)
			tc.mockFunc(client)
			result := client.GetAll(context.Background())
			assert.Equal(t, tc.want, result)
		})
	}
}

func (m *mockGrpc) Update(ctx context.Context, in *pb.ClientRequest, _ ...googlegrpc.CallOption) *pb.ErrorResponse {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.ErrorResponse)
}

func TestService_Update(t *testing.T) {
	cases := []struct {
		name        string
		req         *pb.ClientRequest
		expectedRes interface{}
		want        *pb.ErrorResponse
		mockFunc    func(client *mockGrpc, req interface{})
	}{
		{
			name: "success to update client",
			req: &pb.ClientRequest{
				Name:           "Teste",
				Email:          "manu@gmail.com",
				DocumentNumber: "10025815411",
				PhoneNumber:    "87999991111",
				Cep:            "54406100",
			},
			expectedRes: &pb.ErrorResponse{
				Error: "",
			},
			mockFunc: func(client *mockGrpc, req interface{}) {
				client.On("Update", mock.Anything, req).
					Return(&pb.ErrorResponse{
						Error: "",
					}).Once()
			},
			want: &pb.ErrorResponse{
				Error: "",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := new(mockGrpc)
			tc.mockFunc(client, tc.req)
			result := client.Update(context.Background(), tc.req)
			assert.Equal(t, tc.want, result)
		})
	}
}

func (m *mockGrpc) GetClientByDocumentNumber(ctx context.Context, in *pb.DocNumberRequest, _ ...googlegrpc.CallOption) *pb.ClientResponse {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.ClientResponse)
}

func TestService_GetClientByDocumentNumber(t *testing.T) {
	cases := []struct {
		name        string
		req         *pb.DocNumberRequest
		expectedRes interface{}
		want        *pb.ClientResponse
		mockFunc    func(client *mockGrpc, req interface{})
	}{
		{
			name: "success to get client by document number",
			req: &pb.DocNumberRequest{
				DocumentNumber: "10025815411",
			},
			expectedRes: &pb.ClientResponse{
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
				Error: "",
			},
			mockFunc: func(client *mockGrpc, req interface{}) {
				client.On("GetClientByDocumentNumber", mock.Anything, req).
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
						Error: "",
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
				Error: "",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := new(mockGrpc)
			tc.mockFunc(client, tc.req)
			result := client.GetClientByDocumentNumber(context.Background(), tc.req)
			assert.Equal(t, tc.want, result)
		})
	}
}

func (m *mockGrpc) Delete(ctx context.Context, in *pb.DocNumberRequest, _ ...googlegrpc.CallOption) *pb.ErrorResponse {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.ErrorResponse)
}

func TestService_Delete(t *testing.T) {
	cases := []struct {
		name        string
		req         *pb.DocNumberRequest
		expectedRes interface{}
		want        *pb.ErrorResponse
		mockFunc    func(client *mockGrpc, req interface{})
	}{
		{
			name: "success to delete client",
			req: &pb.DocNumberRequest{
				DocumentNumber: "10025815411",
			},
			expectedRes: &pb.ErrorResponse{
				Error: "",
			},
			mockFunc: func(client *mockGrpc, req interface{}) {
				client.On("Delete", mock.Anything, req).
					Return(&pb.ErrorResponse{
						Error: "",
					}).Once()
			},
			want: &pb.ErrorResponse{
				Error: "",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := new(mockGrpc)
			tc.mockFunc(client, tc.req)
			result := client.Delete(context.Background(), tc.req)
			assert.Equal(t, tc.want, result)
		})
	}
}
