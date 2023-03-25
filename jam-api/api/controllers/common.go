package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paoloposso/jam/libs/core/customerrors"
)

func HandleHttpError(c *gin.Context, err error) {
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
