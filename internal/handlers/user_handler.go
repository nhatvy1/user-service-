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

func (uh *UserHandler) FindUserByID(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := uh.userService.FindUserByID(ctx, 1)

	if err != nil {
		c.JSON(400, gin.H{"error": "Not Found"})
		return
	}

	c.JSON(200, gin.H{"result": user})
}
