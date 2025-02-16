package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Database  DatabaseConfig
	Server    ServerConfig
	JWTSecret string
}

type DatabaseConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type ServerConfig struct {
	Port string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return &Config{
		Database: DatabaseConfig{
			User: getEnv("DATABASE_USER", "postgres"),
			Pass: getEnv("DATABASE_PASSWORD", "password"),
			Host: getEnv("DATABASE_HOST", "db"),
			Port: getEnv("DATABASE_PORT", "5432"),
			Name: getEnv("DATABASE_NAME", "shop"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		JWTSecret: getEnv("JWT_SECRET", "superSecretKey"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
