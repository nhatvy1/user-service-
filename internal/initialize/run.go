package initialize

import (
	"fmt"
	"log"
	di_container "user-service/internal/di-container"
	"user-service/internal/routes"
	"user-service/internal/validations"
	"user-service/pkg/globals"

	"github.com/gin-gonic/gin"
)

func Run() {
	if err := validations.InitValidator(); err != nil {
		fmt.Printf("validations error")
	}

	LoadConfig()

	container := di_container.NewContainer()

	r := gin.Default()

	routes.SetupRoutes(r, container)

	addr := fmt.Sprintf(":%d", globals.Config.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
