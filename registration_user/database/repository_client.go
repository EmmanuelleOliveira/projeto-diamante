package database

import (
	"database/sql"

	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/domain"
)

type ClientRepository struct {
	Db *sql.DB
}

func (c *ClientRepository) Save(client *domain.Client) error {
	stmt, err := c.Db.Prepare("INSERT INTO user(name, email, document_number, phone_number, cep, street, city, uf) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		client.Name,
		client.Email,
		client.DocumentNumber,
		client.PhoneNumber,
		client.Address.Cep,
		client.Address.Street,
		client.Address.City,
		client.Address.UF,
	)
	if err != nil {
		return err
	}

	return nil
}

type ListClients struct {
	Clients []*domain.Client
}

func (c *ClientRepository) GetAll() (*ListClients, error) {
	response, err := c.Db.Query("SELECT * FROM user")

	if err != nil {
		return &ListClients{}, err
	}
	defer response.Close()

	var clients []*domain.Client

	for response.Next() {
		var client domain.Client
		var address domain.AddressClient
		err := response.Scan(
			&client.Id,
			&client.Name,
			&client.Email,
			&client.DocumentNumber,
			&client.PhoneNumber,
			&address.Cep,
			&address.Street,
			&address.City,
			&address.UF,
		)
		if err != nil {
			return nil, err
		}
		client.Address = address
		clients = append(clients, &client)
	}

	return &ListClients{
		Clients: clients,
	}, nil
}

func (c *ClientRepository) GetClientByDocumentNumber(docNumber string) (*domain.Client, error) {
	var client domain.Client
	var address domain.AddressClient

	response := c.Db.QueryRow("SELECT id, name, email, document_number, phone_number, cep, street, city, uf FROM user WHERE document_number = ?", docNumber)

	err := response.Scan(
		&client.Id,
		&client.Name,
		&client.Email,
		&client.DocumentNumber,
		&client.PhoneNumber,
		&address.Cep,
		&address.Street,
		&address.City,
		&address.UF,
	)
	if err != nil {
		return nil, err
	}

	client.Address = address

	return &client, nil
}

func (c *ClientRepository) Update(client *domain.Client) error {
	stmt, err := c.Db.Prepare("UPDATE user SET name=?, email=?, document_number=?, phone_number=?, cep=?, street=?, city=?, uf=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		client.Name,
		client.Email,
		client.DocumentNumber,
		client.PhoneNumber,
		client.Address.Cep,
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

func (c *ClientRepository) Delete(id int) error {
	stmt, err := c.Db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

//Criar func getAll() lista de clients{}
//Criar getClientByDocumentNumber(Doc)
//Criar func update e delete
//select * FROM user
