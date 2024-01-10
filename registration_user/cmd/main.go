package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/client/pb"
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/database"
	"github.com/EmmanuelleOliveira/projeto-diamante/registration_user/domain/client"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 3000, "The port on which the server will listen")
)

func main() {
	db, err := database.ConnectionDB()
	if err != nil {
		fmt.Println("Erro na conexão com o banco de dados:", err)
	}

	service := client.Service{
		Repository:          &database.ClientRepository{Db: db},
		ClientServiceServer: pb.UnimplementedClientServiceServer{},
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterClientServiceServer(grpcServer, service.ClientServiceServer)
	grpcServer.Serve(lis)

}

// var (
// 	port = flag.Int("port", 3000, "The port on which the server will listen")
// )

// func main() {
// 	db, err := database.ConnectionDB()
// 	if err != nil {
// 		fmt.Println("Erro na conexão com o banco de dados:", err)
// 	}

// 	service := &client.Service{
// 		Repository: &database.ClientRepository{Db: db},
// 	}

// 	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	var opts []grpc.ServerOption

// 	grpcServer := grpc.NewServer(opts...)
// 	pb.RegisterClientServiceServer(grpcServer, service)
// 	grpcServer.Serve(lis)
// }
