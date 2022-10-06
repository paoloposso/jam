package main

import (
	"log"
	"net/http"
	"time"

	"github.com/paoloposso/jam/api/config"
	"github.com/paoloposso/jam/api/controllers"
	"github.com/paoloposso/jam/api/src/infrastructure/database"
	"github.com/paoloposso/jam/api/src/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	_ = godotenv.Load()
	gin.SetMode(config.GetGinMode())
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
