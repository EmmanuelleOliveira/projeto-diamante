package main

import (
	"context"
	"fmt"

	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/client/pb"
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/database"
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/domain/client"
)

func main() {
	db, err := database.ConnectionDB()

	if err != nil {
		fmt.Println("Erro na conexão com o banco de dados:", err)
	}

	client1 := &pb.ClientRequest{
		Name:           "Fulano",
		Email:          "fulano@test.com",
		DocumentNumber: "12345678923",
		PhoneNumber:    "1234567890",
		Cep:            "12240150",
	}

	client2 := &pb.ClientRequest{
		Name:           "Siclano",
		Email:          "siclano@test.com",
		DocumentNumber: "12123456835",
		PhoneNumber:    "1234567890",
		Cep:            "12228-030",
	}

	clientToUpdate := &pb.ClientRequest{
		Name:           "Beltrano",
		Email:          "beltrano@test.com",
		DocumentNumber: "12123456845",
		PhoneNumber:    "987654321",
		Cep:            "01017911",
	}

	service := client.Service{
		Repository:          &database.ClientRepository{Db: db},
		ClientServiceServer: pb.UnimplementedClientServiceServer{},
	}

	var ctx context.Context

	//Criação de novos usuários
	i := service.CreateClient(ctx, client1)
	if err != nil {
		fmt.Println("Erro ao salvar o cliente:", err)
	} else {
		fmt.Println("Cliente salvo com sucesso", i)
	}

	x := service.CreateClient(ctx, client2)
	if err != nil {
		fmt.Println("Erro ao salvar o cliente:", err)
	} else {
		fmt.Println("Cliente salvo com sucesso", x)
	}

	//Busca de todos os clientes
	listClients := service.GetAll(ctx)
	if err != nil {
		fmt.Println("Erro ao buscar a lista de clientes:", err)
	} else {
		fmt.Println("Lista de clientes:")
		for _, client := range listClients.Clients {
			fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
				client.Id,
				client.Name,
				client.Email,
				client.DocumentNumber,
				client.PhoneNumber,
				client.Cep,
				client.Address.Street,
				client.Address.City,
				client.Address.Uf,
			)
		}
	}

	// Buscar cliente com número de CPF
	client := service.GetClientByDocumentNumber(ctx, "12345678912")
	if err != nil {
		fmt.Println("Erro ao buscar o cliente:", err)
	} else {
		fmt.Println("Cliente encontrado:")
		fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
			client.Client.Id,
			client.Client.Name,
			client.Client.Email,
			client.Client.DocumentNumber,
			client.Client.PhoneNumber,
			client.Client.Cep,
			client.Client.Address.Street,
			client.Client.Address.City,
			client.Client.Address.Uf,
		)
	}

	// Chamada da função update
	updateClient := service.Update(ctx, clientToUpdate)
	if updateClient != nil {
		fmt.Println("Erro na atualização dos dados do cliente:", err)
	} else {
		fmt.Println("Cliente atualizado com sucesso!")
		clientUpdated := service.GetClientByDocumentNumber(ctx, "12123456879")
		if err != nil {
			fmt.Println("Erro ao buscar o cliente:", err)
		} else {
			fmt.Println("Cliente encontrado:")
			fmt.Printf("ID: %d, Name: %s, Email: %s, Document Number: %s, Phone Number: %s, CEP: %s, Street: %s, City: %s, UF: %s\n",
				clientUpdated.Client.Id,
				clientUpdated.Client.Name,
				clientUpdated.Client.Email,
				clientUpdated.Client.DocumentNumber,
				clientUpdated.Client.PhoneNumber,
				clientUpdated.Client.Cep,
				clientUpdated.Client.Address.Street,
				clientUpdated.Client.Address.City,
				clientUpdated.Client.Address.Uf,
			)
		}
	}

	//Chamada da função delete
	deleteErr := service.Delete(ctx, client2.DocumentNumber)
	if deleteErr != nil {
		fmt.Println("Erro ao excluir dados do cliente:", err)
	} else {
		listClients := service.GetAll(ctx)
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
				client.Cep,
				client.Address.Street,
				client.Address.City,
				client.Address.Uf,
			)
		}

	}
}
