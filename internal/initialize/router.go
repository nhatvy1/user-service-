package initialize

import (
	"user-service/internal/controllers"
	"user-service/internal/routers"
	"user-service/internal/services"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	api := r.Group("/api/v1")

	routers.UserRouter(api, userController)

	return r
}
