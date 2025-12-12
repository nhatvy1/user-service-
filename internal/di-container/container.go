package di_container

import (
	"fmt"
	"user-service/internal/database"
	"user-service/internal/db"
	"user-service/internal/handlers"
	"user-service/internal/services"
	"user-service/pkg/cache"
)

type Container struct {
	Queries *db.Queries
	DB      *database.Database

	UserService services.UserService
	AuthService services.AuthService

	UserHandler *handlers.UserHandler
	AuthHandler *handlers.AuthHandler
}

func NewContainer() *Container {
	// 1. Initialize database
	database, err := database.NewDatabase()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize database: %v", err))
	}

	queries := db.New(database.Pool)

	redis := cache.NewRedisCache()

	userService := services.NewUserService(queries)
	authService := services.NewAuthService(queries, redis)

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	return &Container{
		DB:          database,
		UserService: userService,
		AuthService: authService,
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
}
