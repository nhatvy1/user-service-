package routes

import (
	di_container "user-service/internal/di-container"
	"user-service/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, c *di_container.Container) {
	r.Use(middlewares.Cors())
	r.Use(middlewares.RequestID())
	r.Use(middlewares.Logger())

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api/v1")
	{
		AuthRouter(api, c)
	}
}
