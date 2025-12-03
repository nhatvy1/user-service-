package routers

import (
	"user-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(rg *gin.RouterGroup, uc *controllers.UserController) {
	userGroup := rg.Group("/users")
	{
		userGroup.GET("/:id", uc.FindUserByID)
	}
}
