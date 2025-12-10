package routes

import (
	di_container "user-service/internal/di-container"

	"github.com/gin-gonic/gin"
)

func AuthRouter(rg *gin.RouterGroup, c *di_container.Container) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/login", c.AuthHandler.Login)
		authGroup.POST("/register", c.AuthHandler.Register)
	}
}
