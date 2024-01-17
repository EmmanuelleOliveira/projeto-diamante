package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/client/pb"
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/database"
)

var (
	ErrDocumentNumberInvalid = errors.New("document number invalid")
	ErrClientAlreadyExists   = errors.New("client already exist in our system")
	ErrZipCodeInvalid        = errors.New("zip code invalid")
	ErrSavingClient          = errors.New("error saving client")
	ErrClientNotExist        = errors.New("client not exist")
	ErrUpdateClient          = errors.New("error updating client")
	ErrGetAllCLients         = errors.New("get all clients fail")
	ErrClientDelete          = errors.New("error deleting client")
)

type ClientService struct {
	pb.UnimplementedClientServiceServer
	ClientDB database.Client
}

func NewClientService(clientDb database.Client) *ClientService {
	return &ClientService{
		ClientDB: clientDb,
	}
}

func (s *ClientService) CreateClient(ctx context.Context, client *pb.ClientRequest) (*pb.ClientResponse, error) {
	documentValidated, err := s.ValidateDocumentNumber(client.DocumentNumber)
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  err.Error(),
		}, nil
	}

	client.DocumentNumber = documentValidated

	clientExist, _ := s.ClientDB.CheckClientExists(client.DocumentNumber)
	if clientExist != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrClientAlreadyExists.Error(),
		}, nil
	}

	cep, err := s.ValidateCep(client.Cep)
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrZipCodeInvalid.Error(),
		}, nil
	}

	client.Cep = cep

	addressObj, err := SearchZipCode(cep)
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrZipCodeInvalid.Error(),
		}, nil
	}

	clientObj := database.NewClient(client, addressObj)
	fmt.Println(clientObj)

	err = s.ClientDB.Save(clientObj)
	fmt.Println("Err: ", err)
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrSavingClient.Error(),
		}, nil
	}

	return &pb.ClientResponse{
		Client: &pb.Client{
			Id:             int32(clientObj.Id),
			Name:           clientObj.Name,
			Email:          clientObj.Email,
			DocumentNumber: clientObj.DocumentNumber,
			PhoneNumber:    clientObj.PhoneNumber,
			Cep:            clientObj.Cep,
			Address: &pb.Address{
				Street: clientObj.Address.Street,
				City:   clientObj.Address.City,
				Uf:     clientObj.Address.UF,
			},
		},
		Error: "",
	}, nil
}

func (s *ClientService) Update(ctx context.Context, client *pb.ClientRequest) (*pb.ErrorResponse, error) {
	var addressObj *database.AddressClient

	documentValidated, err := s.ValidateDocumentNumber(client.DocumentNumber)
	if err != nil {
		return &pb.ErrorResponse{
			Error: err.Error(),
		}, err
	}

	client.DocumentNumber = documentValidated

	clientExist, _ := s.ClientDB.CheckClientExists(client.DocumentNumber)
	if clientExist == nil {
		return &pb.ErrorResponse{
			Error: ErrClientNotExist.Error(),
		}, err
	}

	if client.Cep != clientExist.Cep {
		cep, err := s.ValidateCep(client.Cep)
		if err != nil {
			return &pb.ErrorResponse{
				Error: err.Error(),
			}, err
		}

		addressObj, err = SearchZipCode(cep)
		if err != nil {
			return &pb.ErrorResponse{
				Error: ErrZipCodeInvalid.Error(),
			}, err
		}
	} else {
		addressObj = &database.AddressClient{
			Street: clientExist.Address.Street,
			City:   clientExist.Address.City,
			UF:     clientExist.Address.UF,
		}
	}

	clientObj := database.NewClient(client, addressObj)

	clientObj.Id = clientExist.Id

	err = s.ClientDB.Update(clientObj)
	if err != nil {
		return &pb.ErrorResponse{
			Error: ErrUpdateClient.Error(),
		}, err
	}

	return &pb.ErrorResponse{
		Error: "",
	}, nil

}

func (s *ClientService) GetAllClients(ctx context.Context, _ *pb.EmptyField) (*pb.GetAllClientsResponse, error) {
	var clients []*database.Client
	var clientsResponse []*pb.Client

	clients, err := s.ClientDB.GetAll()
	if err != nil {
		return &pb.GetAllClientsResponse{
			Clients: nil,
			Error:   ErrGetAllCLients.Error(),
		}, err
	}

	for _, value := range clients {
		clientObj := &pb.Client{
			Id:             int32(value.Id),
			Name:           value.Name,
			Email:          value.Email,
			DocumentNumber: value.DocumentNumber,
			PhoneNumber:    value.PhoneNumber,
			Cep:            value.Cep,
			Address: &pb.Address{
				Street: value.Address.Street,
				City:   value.Address.City,
				Uf:     value.Address.UF,
			},
		}
		clientsResponse = append(clientsResponse, clientObj)
	}
	return &pb.GetAllClientsResponse{
		Clients: clientsResponse,
		Error:   "",
	}, nil
}

func (s *ClientService) GetClientByDocumentNumber(ctx context.Context, documentNumber *pb.DocNumberRequest) (*pb.ClientResponse, error) {
	documentValidated, err := s.ValidateDocumentNumber(documentNumber.String())
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrDocumentNumberInvalid.Error(),
		}, err
	}

	client, err := s.ClientDB.GetClientByDocumentNumber(documentValidated)
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrClientNotExist.Error(),
		}, err
	}

	return &pb.ClientResponse{
		Client: &pb.Client{
			Id:             int32(client.Id),
			Name:           client.Name,
			Email:          client.Email,
			DocumentNumber: client.DocumentNumber,
			PhoneNumber:    client.PhoneNumber,
			Cep:            client.Cep,
			Address: &pb.Address{
				Street: client.Address.Street,
				City:   client.Address.City,
				Uf:     client.Address.UF,
			},
		},
		Error: "",
	}, nil
}

func (s *ClientService) DeleteClient(ctx context.Context, documentNumber *pb.DocNumberRequest) (*pb.ErrorResponse, error) {
	documentValidated, err := s.ValidateDocumentNumber(documentNumber.String())
	if err != nil {
		return &pb.ErrorResponse{
			Error: ErrDocumentNumberInvalid.Error(),
		}, err
	}

	err = s.ClientDB.Delete(documentValidated)
	if err != nil {
		return &pb.ErrorResponse{
			Error: ErrClientDelete.Error(),
		}, err
	}

	return &pb.ErrorResponse{
		Error: "",
	}, nil

}

func (s *ClientService) ValidateDocumentNumber(documentNumber string) (string, error) {
	documentNumberObj, _ := regexp.Compile(`[-.,-]`)
	docNumber := documentNumberObj.ReplaceAllString(documentNumber, "")

	if len(docNumber) != 11 {
		return "", ErrDocumentNumberInvalid
	}

	return docNumber, nil

}

func (s *ClientService) ValidateCep(cep string) (string, error) {
	cepValidated := regexp.MustCompile(`[-.]`).ReplaceAllString(cep, "")

	if len(cepValidated) == 8 {
		return cepValidated, nil
	}
	return "", ErrZipCodeInvalid

}

func SearchZipCode(cep string) (*database.AddressClient, error) {
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonResponse struct {
		Erro bool `json:"erro"`
	}

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return nil, err
	}

	if jsonResponse.Erro {
		return nil, errors.New("zip code invalid")
	}

	var c database.AddressClient

	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	if c.Street == "" && c.City == "" && c.UF == "" {
		return nil, errors.New("zip code invalid")
	}

	return &c, nil
}
