package main

import (
	"api-graphql/graph"
	"api-graphql/graph/generated"
	"api-graphql/infrastructure/database"
	"api-graphql/users"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	client, err := database.GetClient("mongodb://localhost:27017/jamapp")
	if err != nil {
		panic(err)
	}
	repo := database.UserRepository{Client: client}
	service := users.Service{Repository: repo}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Service: &service}})))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
