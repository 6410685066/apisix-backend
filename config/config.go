package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser       string
	DBPassword   string
	DBHost       string
	DBName       string
	APIPrefix    string
	APIKey       string
	APIPort      string
	AllowOrigins string
}

var AppConfig Config

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not loaded (defaulting to environment variables)")
	}
	AppConfig = Config{
		DBUser:       getEnv("DB_USER", "demo_user"),
		DBPassword:   getEnv("DB_PASSWORD", "demo_pass"),
		DBHost:       getEnv("DB_HOST", "localhost:3306"),
		DBName:       getEnv("DB_NAME", "demo_db"),
		APIPrefix:    getEnv("API_PREFIX", "/api"),
		AllowOrigins: getEnv("ALLOW_ORIGINS", "http://localhost:5173"),
	}

	log.Println("Config loaded:", AppConfig)

	return AppConfig
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func GetJWTSecret() []byte {
	return []byte(AppConfig.APIKey)
}
