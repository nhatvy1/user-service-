package controllers

import (
	"user-service/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(us services.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

func (c *UserController) FindUserByID(ctx *gin.Context) {
	user, err := c.userService.FindUserByID(1)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Not Found"})
		return
	}

	ctx.JSON(200, gin.H{"result": user})
}
