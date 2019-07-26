package main

import (
	anime_service "anime-service"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/nats-io/nats.go"
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

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))

	http.Handle("/query", handler.GraphQL(anime_service.NewExecutableSchema(anime_service.Config{Resolvers: anime_service.NewResolver(nc)})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
