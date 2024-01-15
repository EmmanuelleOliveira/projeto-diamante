package database

import (
	"database/sql"

	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/client/pb"
)

type Client struct {
	db             *sql.DB
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

func NewClientRepository(db *sql.DB) *Client {
	return &Client{db: db}
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

func (c *Client) Save(client *Client) error {
	stmt, err := c.db.Prepare("INSERT INTO user(name, email, document_number, phone_number, cep, street, city, uf) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		client.Name,
		client.Email,
		client.DocumentNumber,
		client.PhoneNumber,
		client.Cep,
		client.Address.Street,
		client.Address.City,
		client.Address.UF,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetAll() ([]*Client, error) {
	response, err := c.db.Query("SELECT * FROM user")

	if err != nil {
		return nil, err
	}
	defer response.Close()

	var clients []*Client

	for response.Next() {
		client := &Client{}

		err := response.Scan(
			&client.Id,
			&client.Name,
			&client.Email,
			&client.DocumentNumber,
			&client.PhoneNumber,
			&client.Cep,
			&client.Address.Street,
			&client.Address.City,
			&client.Address.UF,
		)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func (c *Client) GetClientByDocumentNumber(docNumber string) (*Client, error) {
	stmt, err := c.db.Prepare("SELECT * FROM user WHERE document_number = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	client := &Client{}

	err = stmt.QueryRow(docNumber).Scan(
		&client.Id,
		&client.Name,
		&client.Email,
		&client.DocumentNumber,
		&client.PhoneNumber,
		&client.Cep,
		&client.Address.Street,
		&client.Address.City,
		&client.Address.UF,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Update(client *Client) error {
	stmt, err := c.db.Prepare("UPDATE user SET name=?, email=?, document_number=?, phone_number=?, cep=?, street=?, city=?, uf=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		client.Name,
		client.Email,
		client.DocumentNumber,
		client.PhoneNumber,
		client.Cep,
		client.Address.Street,
		client.Address.City,
		client.Address.UF,
		client.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete(documentNumber string) error {
	stmt, err := c.db.Prepare("DELETE FROM user WHERE document_number=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(documentNumber)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CheckClientExists(docNumber string) (*Client, error) {
	stmt, err := c.db.Prepare("SELECT * FROM user WHERE document_number = ?")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	client := &Client{}

	err = stmt.QueryRow(docNumber).Scan(
		&client.Id,
		&client.Name,
		&client.Email,
		&client.DocumentNumber,
		&client.PhoneNumber,
		&client.Cep,
		&client.Address.Street,
		&client.Address.City,
		&client.Address.UF,
	)

	if err != nil {
		return nil, err
	}

	if client.DocumentNumber != "" {
		return client, nil
	}

	return nil, nil

}
