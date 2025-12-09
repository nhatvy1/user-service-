package routes

import (
	di_container "user-service/internal/di-container"
	"user-service/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(rg *gin.RouterGroup, c *di_container.Container) {
	userGroup := rg.Group("/users")
	{
		protected := userGroup.Group("/")
		protected.Use(middlewares.Authentication())
		{
			protected.GET("/:id", c.UserHandler.FindUserByID)
			protected.PUT("/:id", c.UserHandler.FindUserByID)
			protected.DELETE("/:id", c.UserHandler.FindUserByID)
		}
	}
}
