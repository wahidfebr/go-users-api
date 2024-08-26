package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppHost string
	AppPort string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return Config{
		AppHost: os.Getenv("APP_HOST"),
		AppPort: os.Getenv("APP_PORT"),
	}
}
