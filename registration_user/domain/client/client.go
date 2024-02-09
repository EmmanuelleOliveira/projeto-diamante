package client

import (
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/client/pb"
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/database"
)

func NewClient(client *pb.ClientRequest, address *database.AddressClient) *database.Client {
	return &database.Client{
		Name:           client.Name,
		Email:          client.Email,
		DocumentNumber: client.DocumentNumber,
		PhoneNumber:    client.PhoneNumber,
		Cep:            client.Cep,
		Address: database.AddressClient{
			Street: address.Street,
			City:   address.City,
			UF:     address.UF,
		},
	}
}
