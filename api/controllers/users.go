package controllers

import (
	"api/src/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *users.Service
}

func NewUserController(router *gin.Engine, service users.Service) {
	controller := UserController{service: &service}

	router.GET("/users/:id", controller.getUser)
	router.GET("/users", controller.getUserByEmail)
	router.POST("/users", controller.create)
}

func (controller *UserController) getUser(c *gin.Context) {
	user, err := controller.service.GetById(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (controller *UserController) getUserByEmail(c *gin.Context) {
	if email, ok := c.GetQuery("email"); ok {
		user, err := controller.service.GetById(email)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, user)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "no parameters provided"})
	}
}

func (controller *UserController) create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := controller.service.InsertUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, result)
}
