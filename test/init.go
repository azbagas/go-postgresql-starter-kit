package test

import (
	"go-postgresql-starter-kit/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var app *fiber.App

var db *gorm.DB

var appConfig *config.Config

var log *logrus.Logger

var validate *validator.Validate

func init() {
	appConfig = config.NewConfig()
	log = config.NewLogger(appConfig)
	validate = config.NewValidator(appConfig)
	app = config.NewFiber(appConfig)
	db = config.NewDatabase(appConfig, log)
	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   appConfig,
	})
}
