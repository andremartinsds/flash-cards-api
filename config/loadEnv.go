package config

import (
	"errors"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()

	if err != nil {
		errors.New("Error loading .env file")
	}
	return nil
}
