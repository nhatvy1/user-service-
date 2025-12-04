package initialize

import (
	"user-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	container := BuildContainer()

	api := r.Group("/api/v1")

	routes.UserRouter(api, container.UserHandler)
	routes.AuthRouter(api, container.AuthHandler)

	return r
}
