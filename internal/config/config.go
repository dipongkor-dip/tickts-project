package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	Dsn        string
	JWT_secret string
}

func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:       os.Getenv("port"),
		Dsn:        os.Getenv("dsn"),
		JWT_secret: os.Getenv("JWT_SECRET"),
	}
}
