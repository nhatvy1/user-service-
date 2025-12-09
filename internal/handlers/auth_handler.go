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

func (ah *AuthHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var loginRequest vo.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		validations.HandleValidationError(c, err)
		return
	}

	result, err := ah.authService.Login(ctx, &loginRequest)
	if err != nil {
		utils.ErrorResponse(c, 403, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, 200, "success", result)
}
