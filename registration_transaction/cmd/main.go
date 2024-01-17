package main

import (
	"fmt"

	"github.com/emmanuelleoliveira/projeto-diamante/registration_transaction/database"
)

func main() {
	_, err := database.ConnectionDB()
	if err != nil {
		fmt.Println("Erro na conexão com o banco de dados:", err)
	}
}
