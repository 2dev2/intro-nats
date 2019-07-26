package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/rubiagatra/search-service/pb"
	"github.com/rubiagatra/search-service/service"
	"google.golang.org/grpc"
)

const (
	port = 8000
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterSearchServiceServer(grpcServer, service.NewService())
	log.Printf("Server grpc running on: %d", port)
	grpcServer.Serve(lis)

}
