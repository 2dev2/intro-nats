package main

import (
	"fmt"
	"log"
	"net"

	"github.com/olivere/elastic"
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

	client, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		log.Fatalf("failed to connect elastic: %v", err)
	}

	grpcServer := grpc.NewServer()
	service := service.NewService()
	service.RegisterElasticClient(client)

	pb.RegisterSearchServiceServer(grpcServer, service)
	log.Printf("Server grpc running on: %d", port)
	grpcServer.Serve(lis)

}
