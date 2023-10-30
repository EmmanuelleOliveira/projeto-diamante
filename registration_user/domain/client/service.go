package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
)

var (
	ErrDocumentNumberInvalid = errors.New("document number invalid")
	ErrClientAlreadyExists   = errors.New("client already exist in our system")
	ErrZipCodeInvalid        = errors.New("zip code invalid")
	ErrSavingClient          = errors.New("error saving client")
	ErrClientNotExist        = errors.New("client not exist")
	ErrUpdateClient          = errors.New("error updating client")
	ErrGetAllCLients         = errors.New("get all clients fail")
)

type Service struct {
	Repository Repository
}

func (s *Service) CreateNewClient(client *Client) (*Client, error) {
	documentValidated, err := s.ValidateDocumentNumber(client.DocumentNumber)
	if err != nil {
		return nil, err
	}

	client.DocumentNumber = documentValidated

	clientExist, _ := s.Repository.CheckClientExists(client.DocumentNumber, client.Email)
	if clientExist != nil {
		return nil, ErrClientAlreadyExists
	}

	addressObj, err := SearchZipCode(client.Cep)
	if err != nil {
		return nil, ErrZipCodeInvalid
	}

	clientObj := NewClient(client, addressObj)

	err = s.Repository.Save(clientObj)
	if err != nil {
		return nil, ErrSavingClient
	}

	return clientObj, nil
}

func (s *Service) UpdateClient(client *Client) (*Client, error) {
	var addressObj *AddressClient

	documentValidated, err := s.ValidateDocumentNumber(client.DocumentNumber)
	if err != nil {
		return nil, err
	}

	client.DocumentNumber = documentValidated

	clientExist, _ := s.Repository.CheckClientExists(client.DocumentNumber, client.Email)
	if clientExist == nil {
		return nil, ErrClientNotExist
	}

	if client.Cep != clientExist.Cep {
		addressObj, err = SearchZipCode(client.Cep)
		if err != nil {
			return nil, ErrZipCodeInvalid
		}
	} else {
		addressObj = &AddressClient{
			Street: client.Address.Street,
			City:   client.Address.City,
			UF:     client.Address.UF,
		}
	}

	clientObj := NewClient(client, addressObj)

	err = s.Repository.Update(clientObj)
	if err != nil {
		return nil, ErrUpdateClient
	}

	return clientObj, nil

}

func (s *Service) GetAll() ([]*Client, error) {
	clients, err := s.Repository.GetAll()
	if err != nil {
		return nil, ErrGetAllCLients
	}

	return clients.Clients, nil
}

func (s *Service) GetClientByDocumentNumber(documentNumber string) (*Client, error) {
	documentValidated, err := s.ValidateDocumentNumber(documentNumber)
	if err != nil {
		return nil, ErrDocumentNumberInvalid
	}

	client, err := s.GetClientByDocumentNumber(documentValidated)
	if err != nil {
		return nil, ErrClientNotExist
	}

	return client, nil
}

func (s *Service) ValidateDocumentNumber(documentNumber string) (string, error) {
	documentNumberObj, _ := regexp.Compile(`[-.,-]`)
	docNumber := documentNumberObj.ReplaceAllString(documentNumber, "")

	if len(docNumber) != 11 {
		return "", ErrDocumentNumberInvalid
	}

	return docNumber, nil

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

	var c AddressClient

	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

//criar getAll, getClientByDocumentNumber
