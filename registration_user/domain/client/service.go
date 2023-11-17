package client

import (
	"encoding/json"
	"errors"
	"fmt"
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
	ErrClientDelete          = errors.New("error deleting client")
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

	clientExist, _ := s.Repository.CheckClientExists(client.DocumentNumber)
	if clientExist != nil {
		return nil, ErrClientAlreadyExists
	}

	cep, err := s.ValidateCep(client.Cep)
	if err != nil {
		return nil, err
	}

	client.Cep = cep

	addressObj, err := SearchZipCode(cep)
	if err != nil {
		return nil, ErrZipCodeInvalid
	}

	clientObj := NewClient(client, addressObj)
	fmt.Println(clientObj)

	err = s.Repository.Save(clientObj)
	if err != nil {
		return nil, ErrSavingClient
	}

	return clientObj, nil
}

func (s *Service) Update(client *Client) error {
	var addressObj *AddressClient

	documentValidated, err := s.ValidateDocumentNumber(client.DocumentNumber)
	if err != nil {
		return err
	}

	client.DocumentNumber = documentValidated

	clientExist, _ := s.Repository.CheckClientExists(client.DocumentNumber)
	if clientExist == nil {
		return ErrClientNotExist
	}

	if client.Cep != clientExist.Cep {
		cep, err := s.ValidateCep(client.Cep)
		if err != nil {
			return err
		}

		addressObj, err = SearchZipCode(cep)
		if err != nil {
			return ErrZipCodeInvalid
		}
	} else {
		addressObj = &AddressClient{
			Street: client.Address.Street,
			City:   client.Address.City,
			UF:     client.Address.UF,
		}
	}

	clientObj := NewClient(client, addressObj)

	clientObj.Id = clientExist.Id

	err = s.Repository.Update(clientObj)
	if err != nil {
		return ErrUpdateClient
	}

	return nil

}

func (s *Service) GetAll() ([]*Client, error) {
	var clients []*Client

	clients, err := s.Repository.GetAll()
	if err != nil {
		return nil, ErrGetAllCLients
	}

	return clients, nil
}

func (s *Service) GetClientByDocumentNumber(documentNumber string) (*Client, error) {
	documentValidated, err := s.ValidateDocumentNumber(documentNumber)
	if err != nil {
		return nil, ErrDocumentNumberInvalid
	}

	client, err := s.Repository.GetClientByDocumentNumber(documentValidated)
	if err != nil {
		return nil, ErrClientNotExist
	}

	return client, nil
}

func (s *Service) Delete(documentNumber string) error {
	documentValidated, err := s.ValidateDocumentNumber(documentNumber)
	if err != nil {
		return ErrDocumentNumberInvalid
	}

	err = s.Repository.Delete(documentValidated)
	if err != nil {
		return ErrClientDelete
	}

	return nil

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
