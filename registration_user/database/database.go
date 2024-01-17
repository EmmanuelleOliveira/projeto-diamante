package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() (*sql.DB, error) {
	config := GetConfig()
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Name)

	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	return db, nil
}
