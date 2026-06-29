package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string
	Port   string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("config: no .env file found, using environment variables")
	}

	return &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		Port:   getEnv("PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
	}
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)
}

func (c *Config) IsProduction() bool {
	return c.AppEnv == "production"
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
