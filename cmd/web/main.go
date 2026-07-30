package main

import (
	"fmt"
	"go-postgresql-starter-kit/internal/config"
)

func main() {
	appConfig := config.NewConfig()
	log := config.NewLogger(appConfig)
	db := config.NewDatabase(appConfig, log)
	validate := config.NewValidator(appConfig)
	app := config.NewFiber(appConfig)
	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   appConfig,
	})

	webPort := appConfig.WebPort
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
