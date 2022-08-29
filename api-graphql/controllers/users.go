package controllers

import (
	"api-graphql/src/infrastructure/database"
	"api-graphql/src/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *users.Service
}

func NewUserController(router *gin.Engine, mongoUrl, databaseName string) UserController {
	database := users.NewService(database.NewRepository(mongoUrl, databaseName))

	router.GET("/users/:id", getUser)
	router.GET("/users", getUserByEmail)
	router.POST("/users", create)

	return UserController{service: &database}
}

func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users.User{Name: "Paolo", Email: "paolo@gmail.com"})
}

func getUserByEmail(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users.User{Name: "Paolo", Email: "paolo@gmail.com"})
}

func create(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users.User{Name: "Paolo", Email: "paolo@gmail.com"})
}
