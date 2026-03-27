package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	Port         string
	Environment  string
	JWTSecret    string
	JWTExpiration string
}

var AppConfig *Config

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		DatabaseURL:  getEnv("DATABASE_URL", ""),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", "postgres"),
		DBName:        getEnv("DB_NAME", "pfd_db"),
		Port:          getEnv("PORT", "3000"),
		Environment:   getEnv("ENV", "development"),
		JWTSecret:     getEnv("JWT_SECRET", "your-secret-key-change-this"),
		JWTExpiration: getEnv("JWT_EXPIRATION", "24"),
	}

	AppConfig = cfg
	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) GetDSN() string {
	if c.DatabaseURL != "" {
		return c.DatabaseURL
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
c.DBHost,
c.DBPort,
c.DBUser,
c.DBPassword,
c.DBName,
)
}
