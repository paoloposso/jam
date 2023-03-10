package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paoloposso/jam/src/auth"
	"github.com/paoloposso/jam/src/core/customerrors"
	authrepo "github.com/paoloposso/jam/src/infrastructure/dynamodb/auth"
)

func login(c *gin.Context) {
	var req auth_request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if repo, err := authrepo.NewRepository(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		service := auth.NewService(repo)
		user, err := service.Authenticate(req.Email, req.Password)

		if err != nil {
			handleHttpError(c, err)
			return
		}

		c.IndentedJSON(http.StatusOK, user)
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

func main() {
	router := gin.Default()
	router.POST("/auth", login)

	router.Run("localhost:8080")
}
