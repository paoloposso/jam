package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	authcontroller "github.com/paoloposso/jam/cmd/api/controllers/auth"
	"github.com/paoloposso/jam/cmd/api/docs"
	"github.com/paoloposso/jam/libs/auth"
	authrepo "github.com/paoloposso/jam/libs/infrastructure/dynamodb/auth"
	"github.com/rs/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// HealthCheck
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Health-check
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

// @contact.name   Paolo Posso
// @contact.url    http://www.swagger.io/support
// @contact.email  pvictorsys@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	_ = godotenv.Load(".env")

	gin.SetMode(os.Getenv("GIN_MODE"))

	// credentials file
	// cfg, err := config.LoadDefaultConfig(context.TODO(),
	// 	func(o *config.LoadOptions) error {
	// 		return nil
	// 	})

	keyId := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_ACCESS_KEY_SECRET")
	region := os.Getenv("AWS_REGION")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(keyId, secretKey, "")),
		config.WithRegion(region))

	if err != nil {
		fmt.Printf("error: %v", err)
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	repo, err := authrepo.NewRepository(svc)

	if err != nil {
		panic(err)
	}

	authService := auth.NewService(repo)
	authController := authcontroller.NewAuthController(authService)

	port := "5500"

	docs.SwaggerInfo.Title = "Jam API"
	docs.SwaggerInfo.Description = "Jam API."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "jam.swagger.io"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.POST("/v1/auth", authController.Login)
	router.GET("/v1", HealthCheck)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handler := cors.Default().Handler(router)

	log.Printf("Starting on port %v", port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Printf("Error starting server: %v", err)
	} else {
		log.Printf("Server started on port %v", port)
	}
}
