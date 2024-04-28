package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string `env:"DB_PORT" env-default:"5432"`
	Host      string `env:"DB_HOST" env-default:"localhost"`
	Name      string `env:"DB_NAME" env-default:"demo"`
	User      string `env:"DB_USER" env-default:"postgres"`
	Password  string `env:"DB_PASSWORD" env-default:"postgres"`
	JWTSecret string `env:"DB_JWT_SECRET" env-default:"random"`
}

func InitConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	return &Config{
		Port:      os.Getenv("DB_PORT"),
		Host:      os.Getenv("DB_HOST"),
		Name:      os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASSWORD"),
		JWTSecret: os.Getenv("DB_JWT_SECRET"),
	}
}
