package database

import (
	"database/sql"
)

type TransactionRepository struct {
	Db *sql.DB
}

func (c *TransactionRepository) Save() error {
	return nil
}

func (c *TransactionRepository) GetAll() error {
	return nil
}

func (c *TransactionRepository) GetClientByDocumentNumber() error {
	return nil
}

func (c *TransactionRepository) Update() error {
	return nil
}

func (c *TransactionRepository) Delete() error {
	return nil
}

func (c *TransactionRepository) CheckClientExists() error {
	return nil
}
