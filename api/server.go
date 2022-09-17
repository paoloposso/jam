package main

import (
	"api/config"
	"api/controllers"
	"api/src/infrastructure/database"
	"api/src/users"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	_ = godotenv.Load()
	gin.SetMode(config.GetGinMode())
	router := gin.New()
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
	log.Print("Listening on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
