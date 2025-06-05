package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, falling back to environment variables")
	}

	config := Config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if config.Port == "" || config.DatabaseURL == "" {
		log.Fatal("Missing required environment variables PORT or DATABASE_URL")
	}

	return config
}
