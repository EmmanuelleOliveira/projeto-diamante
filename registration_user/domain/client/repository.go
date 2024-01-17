package client

import (
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/client/pb"
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/database"
)

type Repository interface {
	Save(client *database.Client) error
	GetAll() ([]*database.Client, error)
	GetClientByDocumentNumber(docNumber string) (*database.Client, error)
	Update(client *database.Client) error
	Delete(documentNumber string) error
	CheckClientExists(docNumber string) (*database.Client, error)
	NewClient(client *pb.ClientRequest, address *database.AddressClient) *database.Client
}
