package controllers

import (
	"net/http"
	"task-list/services"
	"task-list/utils"

	"github.com/gin-gonic/gin"
)
type authController struct {
	authService services.AuthService
}

func CreateAuthController(authService services.AuthService) *authController{
	return &authController{
		authService: authService,
	}
}

func(a *authController) GenerateAPIKey(c *gin.Context) {
	apiKey := a.authService.GenerateAPIKey()
    utils.HandleSuccess(c, http.StatusOK, apiKey)
}