package main

import (
	"ip_manager/config"
	"ip_manager/external/http"
	httpHandler "ip_manager/interfaces/http"
	"ip_manager/usecases"
	"log"
)

func main() {
	// Viper Configuration
	cfg := config.LoadConfig()

	manager := usecases.NewIPManager()
	handler := httpHandler.NewHandler(manager)

	//Create root subnet
	_, err := manager.CreateSubnet("10.0.0.0/24", 0)
	if err != nil {
		log.Fatalf("Failed to create root subnet: %v", err)
	}

	//Fiber RESTFull API
	http.StartServer(handler, cfg.Port)
}
