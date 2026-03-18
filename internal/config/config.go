package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	Host        string
	Port        string
	DatabaseDSN string
	APIToken    string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		AppName:     getEnv("APP_NAME", "fiber-rest-api"),
		Host:        getEnv("APP_HOST", ""),
		Port:        getEnv("APP_PORT", "5008"),
		DatabaseDSN: getEnv("DATABASE_DSN", "root:@tcp(127.0.0.1:3306)/go_fiber_1?parseTime=true"),
		APIToken:    getEnv("API_TOKEN", "your-secret-token"),
	}
}

func (c Config) Address() string {
	if c.Host == "" {
		return ":" + c.Port
	}

	return c.Host + ":" + c.Port
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
