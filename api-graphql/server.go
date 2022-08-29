package main

import (
	"api-graphql/config"
	"api-graphql/controllers"
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
	controllers.NewUserController(gin.Default(), mongoUrl, databaseName)
	router.Run("localhost:8080")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
