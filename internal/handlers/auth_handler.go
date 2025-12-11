package handlers

import (
	"user-service/internal/services"
	"user-service/internal/utils"
	"user-service/internal/validations"
	"user-service/internal/vo"
	"user-service/pkg/cache"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
	cache       cache.Cache
}

func NewAuthHandler(as services.AuthService, cache cache.Cache) *AuthHandler {
	return &AuthHandler{
		authService: as,
		cache:       cache,
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
	if utils.HandleError(c, err) {
		return
	}

	utils.SuccessResponse(c, 200, "success", result)
}

func (ah *AuthHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	var registerRequest vo.RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		validations.HandleValidationError(c, err)
		return
	}

	result, err := ah.authService.Register(ctx, &registerRequest)
	if utils.HandleError(c, err) {
		return
	}

	utils.SuccessResponse(c, 201, "user registered successfully", result)
}
