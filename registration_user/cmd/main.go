package main

import (
	"fmt"
	"net"

	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/client/pb"
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/database"
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/domain/client"
	"google.golang.org/grpc"
)

func main() {
	db, err := database.ConnectionDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("connection database ok!")

	clientDB := database.NewClientRepository(db)
	clientService := client.NewClientService(*clientDB)

	grpcServer := grpc.NewServer()
	pb.RegisterClientServiceServer(grpcServer, clientService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
