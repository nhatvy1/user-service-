package initialize

import (
	"fmt"
	"log"
	"strconv"
	"user-service/internal/validations"
	"user-service/pkg/globals"
)

func Run() {
	if err := validations.InitValidator(); err != nil {
		fmt.Printf("validations error")
	}
	LoadConfig()

	pool, err := DbConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	r := InitRouter()
	r.Run(fmt.Sprintf(":%s", strconv.Itoa(globals.Config.Server.Port)))
}
