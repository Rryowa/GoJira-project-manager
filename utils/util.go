package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rryowa/Gojira-project-manager/entity"
)

var Envs = NewConfig()

func NewConfig() *entity.Config {
	err := godotenv.Load("utils/.env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	return &entity.Config{
		Port:      os.Getenv("DB_PORT"),
		Host:      os.Getenv("DB_HOST"),
		Name:      os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASSWORD"),
		JWTSecret: os.Getenv("DB_JWT_SECRET"),
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
