package config

import (
	"errors"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// GetFromEnv reads environment varialbes to get configuration for the OLDapp
func GetConfiguration() (Config, error) {
	// Load variables from .env
	err := godotenv.Load()
	// Check for errors and ignore if file not found
	if err != nil && err.Error() != "open .env: no such file or directory" {
		log.Fatal(err)
	}
	appPort, err := strconv.Atoi("3333")
	if err != nil {
		return Config{}, errors.New("Invalid OLDapp port specified")
	}
	dbPort, err := strconv.Atoi("5432")
	if err != nil {
		return Config{}, errors.New("Invalid jdb port specified")
	}
	appForwardWebpack, err := strconv.ParseBool("false")
	if err != nil {
		return Config{}, errors.New("Invalid OLDapp forward webpack option specified")
	}
	dbSSL, err := strconv.ParseBool("false")
	if err != nil {
		return Config{}, errors.New("Invalid DB SSL option specified")
	}
	config := Config{
		AppPort:           appPort,
		AppSecret:         []byte("secret"),
		AppTheme:          "default",
		AppForwardWebpack: appForwardWebpack,
		AppUserName:       "marcetin",
		AppUserPassword:   []byte("$2y$12$JbJmIwUnnXaaiZCYB8pNm.gt/KMdQ4/alJTrWWLkOasd3iNXZRpEu"),
		DBHost:            "localhost",
		DBPort:            dbPort,
		DBUser:            "marcetin",
		DBPassword:        "",
		DBName:            "okno",
		DBRequireSSL:      dbSSL,
	}
	err = validator.New().Struct(config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
