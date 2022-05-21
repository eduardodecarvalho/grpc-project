package main

import (
	"log"
	"net"

	"github.com/eduardodecarvalho/grpc-project/pb"
	"github.com/eduardodecarvalho/grpc-project/services"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
}
