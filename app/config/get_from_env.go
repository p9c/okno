package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// GetFromEnv reads environment varialbes to get configuration for the app
func GetFromEnv() (Config, error) {
	// Load variables from .env
	err := godotenv.Load()
	// Check for errors and ignore if file not found
	if err != nil && err.Error() != "open .env: no such file or directory" {
		log.Fatal(err)
	}
	appPort, err := strconv.Atoi(os.Getenv("OKNO_APP_PORT"))
	if err != nil {
		return Config{}, errors.New("Invalid app port specified")
	}
	dbPort, err := strconv.Atoi(os.Getenv("OKNO_DB_PORT"))
	if err != nil {
		return Config{}, errors.New("Invalid db port specified")
	}
	appForwardWebpack, err := strconv.ParseBool(os.Getenv("OKNO_APP_FORWARD_WEBPACK"))
	if err != nil {
		return Config{}, errors.New("Invalid app forward webpack option specified")
	}
	dbSSL, err := strconv.ParseBool(os.Getenv("OKNO_DB_REQUIRE_SSL"))
	if err != nil {
		return Config{}, errors.New("Invalid DB SSL option specified")
	}
	config := Config{
		AppPort:           appPort,
		AppSecret:         []byte(os.Getenv("OKNO_APP_SECRET")),
		AppTheme:          os.Getenv("OKNO_APP_THEME"),
		AppForwardWebpack: appForwardWebpack,
		AppUserName:       os.Getenv("OKNO_APP_USER_NAME"),
		AppUserPassword:   []byte(os.Getenv("OKNO_APP_USER_PASS")),
		DBHost:            os.Getenv("OKNO_DB_HOST"),
		DBPort:            dbPort,
		DBUser:            os.Getenv("OKNO_DB_USER"),
		DBPassword:        os.Getenv("OKNO_DB_PASS"),
		DBName:            os.Getenv("OKNO_DB_NAME"),
		DBRequireSSL:      dbSSL,
	}
	err = validator.New().Struct(config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
