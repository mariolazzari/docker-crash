package config

import (
	"context"
	"log"
	"os"
	"strings"
)

type Config struct {
	HTTP ServerConfig
}

type ServerConfig struct {
	Port string
}

func ParseConfig(_ context.Context) (*Config, error) {
	httpServerPort := strings.TrimSpace(os.Getenv("HTTP_SERVER_PORT"))
	if httpServerPort == "" {
		log.Println("Missing 'HTTP_SERVER_PORT' env variable. Using 8080 as default.")
		httpServerPort = "8080"
	}

	result := &Config{
		HTTP: ServerConfig{
			Port: httpServerPort,
		},
	}

	return result, nil
}
