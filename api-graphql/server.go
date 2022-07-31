package main

import (
	"api-graphql/config"
	"api-graphql/graph"
	"api-graphql/graph/generated"
	"api-graphql/infrastructure/database"
	"api-graphql/users"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	_ = godotenv.Load()

	port := config.GetPort()
	if port == "" {
		port = defaultPort
	}

	mongoUrl, mongoDatabase := config.GetMongoUrlAndDatabase()
	service := users.NewService(database.NewRepository(mongoUrl, mongoDatabase))

	http.Handle("/", playground.Handler("GraphQL playground", "/graph"))
	http.Handle("/graph", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Service: &service}})))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
