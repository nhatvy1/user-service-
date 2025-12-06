package initialize

import (
	"fmt"
	"log"
	"user-service/internal/validations"
	"user-service/pkg/globals"

	"github.com/gin-gonic/gin"
)

func Run() {
	// Initialize validator
	if err := validations.InitValidator(); err != nil {
		fmt.Printf("validations error")
	}

	// Load config
	LoadConfig()

	// Initialize container (includes DB connection)
	container, err := NewContainer()
	if err != nil {
		log.Fatal("Failed to initialize container:", err)
	}
	defer container.Close()

	// Initialize Gin engine
	r := gin.Default()

	// Setup routes
	SetupRoutes(r, container)

	// Start server
	addr := fmt.Sprintf(":%d", globals.Config.Server.Port)

	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
