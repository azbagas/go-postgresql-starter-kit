package main

import (
	"fmt"
	"github.com/azbagas/go-postgresql-starter-kit/internal/config"
)

//go:generate swag init -g main.go -d .,../../internal/delivery/http,../../internal/model -o ../../api --ot json,yaml --parseInternal

// @title           Go PostgreSQL Starter Kit
// @version         1.0.0
// @description     Go PostgreSQL Starter Kit API Documentation
// @host            localhost:3000
// @BasePath        /
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description                 Enter authentication token
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
