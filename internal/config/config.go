package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	AppName              string `validate:"required"`
	WebPrefork           bool
	WebPort              int    `validate:"required"`
	LogLevel             int32  `validate:"required"`
	DatabaseUsername     string `validate:"required"`
	DatabasePassword     string
	DatabaseHost         string `validate:"required"`
	DatabasePort         int    `validate:"required"`
	DatabaseName         string `validate:"required"`
	DatabasePoolIdle     int    `validate:"required"`
	DatabasePoolMax      int    `validate:"required"`
	DatabasePoolLifetime int    `validate:"required"`
}

func NewConfig() *Config {
	config := viper.New()

	config.SetConfigName(".env")
	config.SetConfigType("env")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")
	_ = config.ReadInConfig()

	config.AutomaticEnv()

	cfg := Config{
		AppName:              config.GetString("APP_NAME"),
		WebPrefork:           config.GetBool("WEB_PREFORK"),
		WebPort:              config.GetInt("WEB_PORT"),
		LogLevel:             int32(config.GetInt("LOG_LEVEL")),
		DatabaseUsername:     config.GetString("DATABASE_USERNAME"),
		DatabasePassword:     config.GetString("DATABASE_PASSWORD"),
		DatabaseHost:         config.GetString("DATABASE_HOST"),
		DatabasePort:         config.GetInt("DATABASE_PORT"),
		DatabaseName:         config.GetString("DATABASE_NAME"),
		DatabasePoolIdle:     config.GetInt("DATABASE_POOL_IDLE"),
		DatabasePoolMax:      config.GetInt("DATABASE_POOL_MAX"),
		DatabasePoolLifetime: config.GetInt("DATABASE_POOL_LIFETIME"),
	}

	validate := validator.New()
	err := validate.Struct(&cfg)
	if err != nil {
		panic(fmt.Errorf("Fatal error config validation: %w \n", err))
	}

	return &cfg
}
