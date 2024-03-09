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
	ServerConfig   ServerConfig   `env:", prefix=SERVER_"`
}

type ServerConfig struct {
	ADDR string `env:"ADDR" default:":8080"`
}

type DatabaseConfig struct {
	Host     string `env:"HOST" default:"localhost"`
	Port     int    `env:"PORT" default:"5432"`
	Database string `env:"NAME" default:"root"`
	Schema   string `env:"SCHEMA" default:"public"`
	Username string `env:"USERNAME" default:"root"`
	Password string `env:"PASSWORD" default:"root"`
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
