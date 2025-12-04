package routes

import (
	"user-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UserRouter(rg *gin.RouterGroup, uc *handlers.UserHandler) {
	userGroup := rg.Group("/users")
	{
		userGroup.GET("/:id", uc.FindUserByID)
	}
}
