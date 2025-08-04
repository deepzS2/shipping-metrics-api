package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIPort           string
	FreteRapidoAPIUrl string
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBSslMode         string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		APIPort:           getEnv("API_PORT", "8080"),
		DBHost:            getEnv("DB_HOST", "db"),
		DBPort:            getEnv("DB_PORT", "5432"),
		DBUser:            getEnv("DB_USER", "postgres"),
		DBPassword:        getEnv("DB_PASSWORD", "postgres"),
		DBName:            getEnv("DB_NAME", "freterapido"),
		DBSslMode:         getEnv("DB_SSL_MODE", "disable"),
		FreteRapidoAPIUrl: getEnv("FRETE_RAPIDO_API_URL", ""),
	}

	return cfg, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSslMode)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
