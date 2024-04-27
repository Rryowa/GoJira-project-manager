package main

import (
	"fmt"
	"os"
)

//hold env variables at the runtime from injection

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"),
			getEnv("DB_PORT", "3306")),
		DBName:    getEnv("DB_NAME", "ProjectManager"),
		JWTSecret: getEnv("JWT_SECRET", "randomjwtsecretkey"),
	}
}

// Func taht gets var from env
func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}