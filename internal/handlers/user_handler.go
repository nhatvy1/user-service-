package handlers

import (
	"user-service/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(us services.UserService) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (c *UserHandler) FindUserByID(ctx *gin.Context) {
	user, err := c.userService.FindUserByID(ctx, 1)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Not Found"})
		return
	}

	ctx.JSON(200, gin.H{"result": user})
}
