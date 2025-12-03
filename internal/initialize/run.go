package initialize

import (
	"fmt"
	"log"
	"strconv"
	"user-service/pkg/globals"
)

func Run() {
	LoadConfig()

	pool, err := DbConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	r := InitRouter()
	r.Run(fmt.Sprintf(":%s", strconv.Itoa(globals.Config.Server.Port)))
}
