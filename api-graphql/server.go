package main

import (
	"api-graphql/config"
	"api-graphql/controllers"
	"api-graphql/src/infrastructure/database"
	"api-graphql/src/users"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	_ = godotenv.Load()
	router := gin.Default()
	log.Print("Starting server")
	port := config.GetPort()
	if port == "" {
		port = defaultPort
	}
	mongoUrl, databaseName := config.GetMongoUrlAndDatabase()
	controllers.NewUserController(
		router,
		*users.NewService(database.NewUserRepository(mongoUrl, databaseName)),
	)
	router.Run("localhost:8080")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
