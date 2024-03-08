package config

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"log"
	"os"
)

type Config struct {
	DatabaseConfig DatabaseConfig `env:", prefix=DB_"`
}

type DatabaseConfig struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	Database string `env:"NAME"`
	Schema   string `env:"SCHEMA"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
}

func LoadConfig() Config {
	env := os.Getenv("ENV")

	envFile := ".env"

	if env != "" {
		envFile += "." + env
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var c = Config{}

	if err := envconfig.Process(context.TODO(), &c); err != nil {
		log.Fatal(err)
	}

	return c
}
