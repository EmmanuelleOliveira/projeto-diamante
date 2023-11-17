package database

import (
	"database/sql"

	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/domain/client"
)

type ClientRepository struct {
	Db *sql.DB
}

func (c *ClientRepository) Save(client *client.Client) error {
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

func (c *ClientRepository) GetAll() ([]*client.Client, error) {
	response, err := c.Db.Query("SELECT * FROM user")

	if err != nil {
		return nil, err
	}
	defer response.Close()

	var clients []*client.Client

	for response.Next() {
		client := &client.Client{}

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

func (c *ClientRepository) GetClientByDocumentNumber(docNumber string) (*client.Client, error) {
	stmt, err := c.Db.Prepare("SELECT * FROM user WHERE document_number = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	client := &client.Client{}

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

func (c *ClientRepository) Update(client *client.Client) error {
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

func (c *ClientRepository) Delete(documentNumber string) error {
	stmt, err := c.Db.Prepare("DELETE FROM user WHERE document_number=?")
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

func (c *ClientRepository) CheckClientExists(docNumber string) (*client.Client, error) {
	stmt, err := c.Db.Prepare("SELECT * FROM user WHERE document_number = ?")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	client := &client.Client{}

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

//Criar func getAll() lista de clients{}
//Criar getClientByDocumentNumber(Doc)
//Criar func update e delete
//select * FROM user
