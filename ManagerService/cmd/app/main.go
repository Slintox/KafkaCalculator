package main

import (
	"log"

	"ManagerService/config"
	"ManagerService/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.New("configs/config.yaml")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
