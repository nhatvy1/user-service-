package initialize

import (
	"user-service/internal/handlers"
	"user-service/internal/services"
)

type Container struct {
	AuthHandler *handlers.AuthHandler
	UserHandler *handlers.UserHandler
}

func BuildContainer() *Container {

	userService := services.NewUserService()
	userHandler := handlers.NewUserHandler(userService)

	authService := services.NewAuthService()
	authHandler := handlers.NewAuthHandler(authService)

	return &Container{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
}
