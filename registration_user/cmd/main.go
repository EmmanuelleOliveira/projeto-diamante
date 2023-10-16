package main

import (
	"fmt"

	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/database"
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/domain"
)

func main() {
	db, err := database.ConnectionDB()

	if err != nil {
		fmt.Println("Erro na conexão com o banco de dados:", err)
	}

	clientRepo := &database.ClientRepository{Db: db}

	client1 := &domain.Client{
		Name:           "Fulano",
		Email:          "fulano@test.com",
		DocumentNumber: "123456789",
		PhoneNumber:    "1234567890",
		Address: domain.AddressClient{
			Cep:    "12345678",
			Street: "Rua teste",
			City:   "Cidade",
			UF:     "UF",
		},
	}

	client2 := &domain.Client{
		Name:           "Siclano",
		Email:          "siclano@test.com",
		DocumentNumber: "123456879",
		PhoneNumber:    "1234567890",
		Address: domain.AddressClient{
			Cep:    "12345678",
			Street: "Rua",
			City:   "Cidade",
			UF:     "UF",
		},
	}

	err = clientRepo.Save(client1)
	if err != nil {
		fmt.Println("Erro ao salvar o cliente:", err)
	} else {
		fmt.Println("Cliente salvo com sucesso")
	}

	err = clientRepo.Save(client2)
	if err != nil {
		fmt.Println("Erro ao salvar o cliente:", err)
	} else {
		fmt.Println("Cliente salvo com sucesso")
	}

	listClients, err := clientRepo.GetAll()
	if err != nil {
		fmt.Println("Erro ao buscar a lista de clientes:", err)
	}

	fmt.Println("Lista de clientes:")
	for _, client := range listClients.Clients {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
			client.Id,
			client.Name,
			client.Email,
			client.DocumentNumber,
			client.PhoneNumber,
			client.Address.Cep,
			client.Address.Street,
			client.Address.City,
			client.Address.UF,
		)
	}

	client, err := clientRepo.GetClientByDocumentNumber("123456779")
	if err != nil {
		fmt.Println("Erro ao buscar o cliente:", err)
	}

	fmt.Println("Cliente encontrado:")
	if client != nil {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
			client.Id,
			client.Name,
			client.Email,
			client.DocumentNumber,
			client.PhoneNumber,
			client.Address.Cep,
			client.Address.Street,
			client.Address.City,
			client.Address.UF,
		)
	} else {
		fmt.Println("Nenhum cliente encontrado com o número do documento fornecido.")
	}

	clientToUpdate := &domain.Client{
		Id:             1, // Substitua pelo ID correto do cliente que deseja atualizar
		Name:           "Beltrano",
		Email:          "beltrano@test.com",
		DocumentNumber: "123456879",
		PhoneNumber:    "987654321",
		Address: domain.AddressClient{
			Cep:    "12345678",
			Street: "Nova Rua",
			City:   "Nova Cidade",
			UF:     "PE",
		},
	}

	// Chamada da função update
	updateErr := clientRepo.Update(clientToUpdate)
	if updateErr != nil {
		fmt.Println("Erro na atualização dos dados do cliente:", err)
	} else {
		fmt.Println("Cliente atualizado com sucesso!")
		client, err := clientRepo.GetClientByDocumentNumber("123456879")
		if err != nil {
			fmt.Println("Erro ao buscar o cliente:", err)
		}

		fmt.Println("Cliente encontrado:")
		if client != nil {
			fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
				client.Id,
				client.Name,
				client.Email,
				client.DocumentNumber,
				client.PhoneNumber,
				client.Address.Cep,
				client.Address.Street,
				client.Address.City,
				client.Address.UF,
			)
		} else {
			fmt.Println("Nenhum cliente encontrado com o número do documento fornecido.")
		}
	}

	clientToDelete := 1

	deleteErr := clientRepo.Delete(clientToDelete)
	if deleteErr != nil {
		fmt.Println("Erro ao excluir dados do cliente:", err)
	} else {
		listClients, err := clientRepo.GetAll()
		if err != nil {
			fmt.Println("Erro ao buscar a lista de clientes:", err)
		}

		fmt.Println("Lista de clientes atualizada:")
		for _, client := range listClients.Clients {
			fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
				client.Id,
				client.Name,
				client.Email,
				client.DocumentNumber,
				client.PhoneNumber,
				client.Address.Cep,
				client.Address.Street,
				client.Address.City,
				client.Address.UF,
			)
		}

	}
}
