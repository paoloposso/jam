package controllers

import (
	"net/http"

	"github.com/paoloposso/jam/api/src/users"

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
		handleHttpError(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *UserController) getUserByEmail(c *gin.Context) {
	if email, ok := c.GetQuery("email"); ok {
		user, err := controller.service.GetById(email)

		if err != nil {
			handleHttpError(err, c)
			return
		}
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no parameters provided"})
	}
}

func (controller *UserController) create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		handleHttpError(err, c)
		return
	}
	result, err := controller.service.InsertUser(user)
	if err != nil {
		handleHttpError(err, c)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func handleHttpError(err error, c *gin.Context) {
	code, message := GetHttpError(err)
	c.JSON(code, gin.H{"error": message})
}
