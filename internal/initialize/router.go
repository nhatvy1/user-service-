package initialize

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, container *Container) {
	// Public routes
	api := r.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", container.AuthHandler.Login)
		}

		// Protected routes
		protected := api.Group("")
		// protected.Use(middleware.AuthMiddleware())
		{
			users := protected.Group("/users")
			{
				users.GET("/:id", container.UserHandler.FindUserByID)
			}
		}
	}
}
