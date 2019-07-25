package main

import (
	anime_service "anime-service"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))

	http.Handle("/query", handler.GraphQL(anime_service.NewExecutableSchema(anime_service.Config{Resolvers: anime_service.NewResolver()})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
