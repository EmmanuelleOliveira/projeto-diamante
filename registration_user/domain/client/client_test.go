package client

import (
	"testing"

	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/client/pb"
	"github.com/stretchr/testify/assert"
)

var (
	clientMock = &pb.ClientRequest{
		Name:           "Manu",
		Email:          "manu@gmail.com",
		DocumentNumber: "10025815411",
		PhoneNumber:    "87999991111",
		Cep:            "54406100",
	}

	addressMock = &AddressClient{
		Street: "street test",
		City:   "city test",
		UF:     "Test",
	}
)

func TestNewClient(t *testing.T) {
	cases := []struct {
		name   string
		reqOne *pb.ClientRequest
		reqTwo *AddressClient
		want   *Client
	}{
		{
			name:   "success new client",
			reqOne: clientMock,
			reqTwo: addressMock,
			want: &Client{
				Name:           "Manu",
				Email:          "manu@gmail.com",
				DocumentNumber: "10025815411",
				PhoneNumber:    "87999991111",
				Cep:            "54406100",
				Address: AddressClient{
					Street: "street test",
					City:   "city test",
					UF:     "Test",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewClient(tc.reqOne, tc.reqTwo)
			assert.Equal(t, tc.want, result)
		})
	}
}
