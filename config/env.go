package config

import (
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string

	Port     string
	Host string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		DBUser:     getEnv("DB_USER", "dipeshpaneru"),
		DBPassword: getEnv("DB_PASSWORD", "go_ecommerce_123"),
		DBAddress:  getEnv("DB_ADDRESS", "http://localhost:5432"),
		DBName:     getEnv("DB_NAME", "go_ecommerce"),

		Port:     getEnv("PORT", "5432"),
		Host: getEnv("PUBLIC_HOST", "localhost"),
	}
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}