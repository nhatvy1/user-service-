package routes

import (
	"user-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(rg *gin.RouterGroup, uc *handlers.AuthHandler) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/login", uc.Login)
	}
}
