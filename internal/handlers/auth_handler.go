package handlers

import (
	"user-service/internal/services"
	"user-service/internal/utils"
	"user-service/internal/validations"
	"user-service/internal/vo"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(as services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: as,
	}
}

func (c *AuthHandler) Login(ctx *gin.Context) {
	var loginRequest vo.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		validations.HandleValidationError(ctx, err)
		return
	}
	result, err := c.authService.Login()
	if err != nil {
		utils.ErrorResponse(ctx, 403, "Login Failed", err)
		return
	}

	utils.SuccessResponse(ctx, 200, "success", result)
}
