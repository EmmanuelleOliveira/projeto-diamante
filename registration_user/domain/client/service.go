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

type Service struct {
	Repository          Repository
	ClientServiceServer pb.ClientServiceServer
}

func (s *Service) CreateClient(ctx context.Context, client *pb.ClientRequest) *pb.ClientResponse {
	documentValidated, err := s.ValidateDocumentNumber(client.DocumentNumber)
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  err.Error(),
		}
	}

	client.DocumentNumber = documentValidated

	clientExist, _ := s.Repository.CheckClientExists(client.DocumentNumber)
	if clientExist != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrClientAlreadyExists.Error(),
		}
	}

	cep, err := s.ValidateCep(client.Cep)
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrZipCodeInvalid.Error(),
		}
	}

	client.Cep = cep

	addressObj, err := SearchZipCode(cep)
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrZipCodeInvalid.Error(),
		}
	}

	clientObj := NewClient(client, addressObj)
	fmt.Println(clientObj)

	err = s.Repository.Save(clientObj)
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrSavingClient.Error(),
		}
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
	}
}

func (s *Service) Update(ctx context.Context, client *pb.ClientRequest) *pb.ErrorResponse {
	var addressObj *AddressClient

	documentValidated, err := s.ValidateDocumentNumber(client.DocumentNumber)
	if err != nil {
		return &pb.ErrorResponse{
			Error: err.Error(),
		}
	}

	client.DocumentNumber = documentValidated

	clientExist, _ := s.Repository.CheckClientExists(client.DocumentNumber)
	if clientExist == nil {
		return &pb.ErrorResponse{
			Error: ErrClientNotExist.Error(),
		}
	}

	if client.Cep != clientExist.Cep {
		cep, err := s.ValidateCep(client.Cep)
		if err != nil {
			return &pb.ErrorResponse{
				Error: err.Error(),
			}
		}

		addressObj, err = SearchZipCode(cep)
		if err != nil {
			return &pb.ErrorResponse{
				Error: ErrZipCodeInvalid.Error(),
			}
		}
	} else {
		addressObj = &AddressClient{
			Street: clientExist.Address.Street,
			City:   clientExist.Address.City,
			UF:     clientExist.Address.UF,
		}
	}

	clientObj := NewClient(client, addressObj)

	clientObj.Id = clientExist.Id

	err = s.Repository.Update(clientObj)
	if err != nil {
		return &pb.ErrorResponse{
			Error: ErrUpdateClient.Error(),
		}
	}

	return &pb.ErrorResponse{
		Error: "",
	}

}

func (s *Service) GetAll(ctx context.Context) *pb.GetAllClientsResponse {
	var clients []*Client
	var clientsResponse []*pb.Client

	clients, err := s.Repository.GetAll()
	if err != nil {
		return &pb.GetAllClientsResponse{
			Clients: nil,
			Error:   ErrGetAllCLients.Error(),
		}
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
	}
}

func (s *Service) GetClientByDocumentNumber(ctx context.Context, documentNumber string) *pb.ClientResponse {
	documentValidated, err := s.ValidateDocumentNumber(documentNumber)
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrDocumentNumberInvalid.Error(),
		}
	}

	client, err := s.Repository.GetClientByDocumentNumber(documentValidated)
	if err != nil {
		return &pb.ClientResponse{
			Client: nil,
			Error:  ErrClientNotExist.Error(),
		}
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
	}
}

func (s *Service) Delete(ctx context.Context, documentNumber string) *pb.ErrorResponse {
	documentValidated, err := s.ValidateDocumentNumber(documentNumber)
	if err != nil {
		return &pb.ErrorResponse{
			Error: ErrDocumentNumberInvalid.Error(),
		}
	}

	err = s.Repository.Delete(documentValidated)
	if err != nil {
		return &pb.ErrorResponse{
			Error: ErrClientDelete.Error(),
		}
	}

	return &pb.ErrorResponse{
		Error: "",
	}

}

func (s *Service) ValidateDocumentNumber(documentNumber string) (string, error) {
	documentNumberObj, _ := regexp.Compile(`[-.,-]`)
	docNumber := documentNumberObj.ReplaceAllString(documentNumber, "")

	if len(docNumber) != 11 {
		return "", ErrDocumentNumberInvalid
	}

	return docNumber, nil

}

func (s *Service) ValidateCep(cep string) (string, error) {

	re := regexp.MustCompile(`^\d{5}-\d{3}$`)

	if re.MatchString(cep) {
		cepObj, _ := regexp.Compile(`[-.,-]`)
		cepValidated := cepObj.ReplaceAllString(cep, "")
		return cepValidated, nil
	} else {
		return "", ErrZipCodeInvalid
	}
}

func SearchZipCode(cep string) (*AddressClient, error) {
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

	var c AddressClient

	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	if c.Street == "" && c.City == "" && c.UF == "" {
		return nil, errors.New("zip code invalid")
	}

	return &c, nil
}

//criar getAll, getClientByDocumentNumber
