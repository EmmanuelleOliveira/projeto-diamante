package main

import (
	"fmt"

	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/database"
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/domain/client"
)

func main() {
	db, err := database.ConnectionDB()

	if err != nil {
		fmt.Println("Erro na conexão com o banco de dados:", err)
	}

	client1 := &client.Client{
		Name:           "Fulano",
		Email:          "fulano@test.com",
		DocumentNumber: "12345678912",
		PhoneNumber:    "1234567890",
		Cep:            "08596270",
	}

	client2 := &client.Client{
		Name:           "Siclano",
		Email:          "siclano@test.com",
		DocumentNumber: "12123456879",
		PhoneNumber:    "1234567890",
		Cep:            "08596270",
	}

	service := client.Service{
		Repository: &database.ClientRepository{Db: db},
	}

	//Criação de novos usuários
	i, err := service.CreateNewClient(client1)
	if err != nil {
		fmt.Println("Erro ao salvar o cliente:", err)
	} else {
		fmt.Println("Cliente salvo com sucesso", i)
	}

	x, err := service.CreateNewClient(client2)
	if err != nil {
		fmt.Println("Erro ao salvar o cliente:", err)
	} else {
		fmt.Println("Cliente salvo com sucesso", x)
	}

	//Busca de todos os clientes
	listClients, err := service.GetAll()
	if err != nil {
		fmt.Println("Erro ao buscar a lista de clientes:", err)
	} else {
		fmt.Println("Lista de clientes:")
		for _, client := range listClients {
			fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
				client.Id,
				client.Name,
				client.Email,
				client.DocumentNumber,
				client.PhoneNumber,
				client.Cep,
				client.Address.Street,
				client.Address.City,
				client.Address.UF,
			)
		}
	}

	// Buscar cliente com número de CPF
	client, err := service.GetClientByDocumentNumber("12345678912")
	if err != nil {
		fmt.Println("Erro ao buscar o cliente:", err)
	} else {
		fmt.Println("Cliente encontrado:")
		fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
			client.Id,
			client.Name,
			client.Email,
			client.DocumentNumber,
			client.PhoneNumber,
			client.Cep,
			client.Address.Street,
			client.Address.City,
			client.Address.UF,
		)
	}

	// clientToUpdate := &client.Client{
	// 	Id:             1,
	// 	Name:           "Beltrano",
	// 	Email:          "beltrano@test.com",
	// 	DocumentNumber: "12123456879",
	// 	PhoneNumber:    "987654321",
	// 	Cep:            "03985120",
	// }

	// // Chamada da função update
	// updateClient := service.Update(clientToUpdate)
	// if updateClient != nil {
	// 	fmt.Println("Erro na atualização dos dados do cliente:", err)
	// } else {
	// 	fmt.Println("Cliente atualizado com sucesso!")
	// 	client, err := service.GetClientByDocumentNumber("12123456879")
	// 	if err != nil {
	// 		fmt.Println("Erro ao buscar o cliente:", err)
	// 	} else {
	// 		fmt.Println("Cliente encontrado:")
	// 		fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
	// 			client.Id,
	// 			client.Name,
	// 			client.Email,
	// 			client.DocumentNumber,
	// 			client.PhoneNumber,
	// 			client.Cep,
	// 			client.Address.Street,
	// 			client.Address.City,
	// 			client.Address.UF,
	// 		)
	// 	}
	// }

	// 	fmt.Println("Cliente encontrado:")
	// 	if client != nil {
	// 		fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
	// 			client.Id,
	// 			client.Name,
	// 			client.Email,
	// 			client.DocumentNumber,
	// 			client.PhoneNumber,
	// 			client.Cep,
	// 			client.Address.Street,
	// 			client.Address.City,
	// 			client.Address.UF,
	// 		)
	// 	} else {
	// 		fmt.Println("Nenhum cliente encontrado com o número do documento fornecido.")
	// 	}
	// }

	// deleteErr := clientRepo.Delete(client1.DocumentNumber)
	// if deleteErr != nil {
	// 	fmt.Println("Erro ao excluir dados do cliente:", err)
	// } else {
	// 	listClients, err := clientRepo.GetAll()
	// 	if err != nil {
	// 		fmt.Println("Erro ao buscar a lista de clientes:", err)
	// 	}

	// 	fmt.Println("Lista de clientes atualizada:")
	// 	for _, client := range listClients {
	// 		fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
	// 			client.Id,
	// 			client.Name,
	// 			client.Email,
	// 			client.DocumentNumber,
	// 			client.PhoneNumber,
	// 			client.Cep,
	// 			client.Address.Street,
	// 			client.Address.City,
	// 			client.Address.UF,
	// 		)
	// 	}

	// }
}
