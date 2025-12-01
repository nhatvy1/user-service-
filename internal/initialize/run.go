package initialize

import (
	"fmt"
)

func Run() {
	cfg, _ := LoadConfig()

	fmt.Printf("Server: %s, Port: %d \n", cfg.Server.Name, cfg.Server.Port)
}
