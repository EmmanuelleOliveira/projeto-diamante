package client

import "github.com/EmmanuelleOliveira/projeto-diamante/registration_user/client/pb"

type Client struct {
	Id             int
	Name           string
	Email          string
	DocumentNumber string
	PhoneNumber    string
	Cep            string
	Address        AddressClient
}

type AddressClient struct {
	Street string `json:"logradouro"`
	City   string `json:"localidade"`
	UF     string `json:"uf"`
}

func NewClient(client *pb.ClientRequest, address *AddressClient) *Client {
	return &Client{
		Name:           client.Name,
		Email:          client.Email,
		DocumentNumber: client.DocumentNumber,
		PhoneNumber:    client.PhoneNumber,
		Cep:            client.Cep,
		Address: AddressClient{
			Street: address.Street,
			City:   address.City,
			UF:     address.UF,
		},
	}
}
