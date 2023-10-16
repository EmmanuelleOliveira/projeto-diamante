package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() (*sql.DB, error) {
	config := GetConfig()
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Name)
	fmt.Println(url)

	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//Criar docker compose com banco mySQL
// docker compose up -> Subir o container
// docker compose up -d -> Sobe o container e libera o terminal.
// docker compose ls -> Verificar se o container está rodando.
// docker compose exec mysql bash -> Abri um terminal dentro do container.
// mysql -uroot -p diamante (db utilizado / usuario / senha / nome_database)
