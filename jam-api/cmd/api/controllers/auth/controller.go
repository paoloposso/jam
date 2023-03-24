package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	handlers "github.com/paoloposso/jam/cmd/api/controllers"
	"github.com/paoloposso/jam/libs/auth"
)

type AuthController struct {
	service auth.Service
}

func NewAuthController(service auth.Service) AuthController {
	return AuthController{service}
}

// Login
// @Summary User Authentication.
// @Description Performs User Authentication.
// @Tags Login
// @Param data body controllers.AuthRequest true "login data"
// @Produce json
// @Success 200 {object} error
// @Failure 404 {object} error
// @Failure 403 {object} error
// @Failure 500 {object} error
// @Router /auth [post]
func (controller AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := controller.service.Authenticate(req.Email, req.Password)

	if err != nil {
		handlers.HandleHttpError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, LoginResponse{Token: user.Token})
}
