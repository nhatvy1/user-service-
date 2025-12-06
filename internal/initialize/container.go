package initialize

import (
	"user-service/internal/db"
	"user-service/internal/handlers"
	"user-service/internal/services"
)

type Container struct {
	DB *Database

	UserService services.UserService
	AuthService services.AuthService

	UserHandler *handlers.UserHandler
	AuthHandler *handlers.AuthHandler
}

func NewContainer() (*Container, error) {
	// Initialize database
	database, err := NewDatabase()
	if err != nil {
		return nil, err
	}

	queries := db.New(database.Pool)

	userService := services.NewUserService(queries)
	authService := services.NewAuthService()

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	return &Container{
		DB:          database,
		UserService: userService,
		AuthService: authService,
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
