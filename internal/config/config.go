package config

import (
	"os"
)

// Config represents the configuration for the application.
type Config struct {
	Port string
}

// LoadConfig loads the application configuration.
func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port: port,
	}, nil
}
