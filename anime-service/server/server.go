package main

import (
	anime_service "anime-service"
	"log"
	"net/http"
	"os"

	pb "anime-service/pb"

	"github.com/99designs/gqlgen/handler"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(anime_service.NewExecutableSchema(anime_service.Config{Resolvers: anime_service.NewResolver(nc, client)})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
