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
	PublicHost string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "go_ecommerce_123"),
		DBAddress:  getEnv("DB_ADDRESS", "http://localhost:3306"),
		DBName:     getEnv("DB_NAME", "Go_Ecommerce"),
		
		Port:     getEnv("PORT", "8080"),
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
	}
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}