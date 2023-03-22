package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/paoloposso/jam/cmd/api/docs"
	"github.com/paoloposso/jam/cmd/api/handlers"
	"github.com/paoloposso/jam/libs/auth"
	"github.com/paoloposso/jam/libs/core/customerrors"
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

// Login
// @Summary User Authentication.
// @Description Performs User Authentication.
// @Tags Login
// @Param data body handlers.AuthRequest true "login data"
// @Produce json
// @Success 200 {object} error
// @Failure 404 {object} error
// @Failure 403 {object} error
// @Failure 500 {object} error
// @Router /auth [post]
func LoginHandler(c *gin.Context) {
	var req handlers.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	if repo, err := authrepo.NewRepository(svc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		service := auth.NewService(repo)
		user, err := service.Authenticate(req.Email, req.Password)

		if err != nil {
			handleHttpError(c, err)
			return
		}
		c.IndentedJSON(http.StatusOK, handlers.AuthResponse{Token: user.Token})
	}
}

func handleHttpError(c *gin.Context, err error) {
	e := gin.H{"error": err.Error()}

	switch err.(type) {
	case *customerrors.UnauthorizedError:
		c.IndentedJSON(http.StatusUnauthorized, e)
		return
	default:
		c.IndentedJSON(http.StatusInternalServerError, e)
		return
	}
}

// @contact.name   Paolo Posso
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	port := "5500"

	docs.SwaggerInfo.Title = "Jam API"
	docs.SwaggerInfo.Description = "Jam API."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "jam.swagger.io"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.POST("/v1/auth", LoginHandler)
	router.GET("/v1", HealthCheck)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handler := cors.Default().Handler(router)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	} else {
		log.Printf("Server started on port %v", port)
	}
}
