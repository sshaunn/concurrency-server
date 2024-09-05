package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var config *Config

type Config struct {
	ServerPort            string
	EventingToolsBaseURL  string
	EventingToolsEndpoint string
	DownstreamBaseURL     string
	DownstreamEndpoint    string
}

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting current directory: %v", err)
	}
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "local"
	}
	envPath := filepath.Join(currentDir, "config", "resource", ".env."+env)

	log.Printf("Loading environment from: %v", envPath)

	err = godotenv.Load(envPath)
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
		// Continue execution even if .env file is not found
	}

	config = &Config{
		ServerPort:            os.Getenv("SERVER_PORT"),
		EventingToolsBaseURL:  os.Getenv("SERVICE_EVENTING_TOOLS_BASE_URL"),
		EventingToolsEndpoint: os.Getenv("SERVICE_EVENTING_TOOLS_ENDPOINT"),
		DownstreamBaseURL:     os.Getenv("SERVICE_DOWNSTREAM_BASE_URL"),
		DownstreamEndpoint:    os.Getenv("SERVICE_DOWNSTREAM_ENDPOINT"),
	}

	// Validate required fields
	if config.ServerPort == "" {
		log.Fatal("SERVER_PORT is not set")
	}
}

func Load() *Config {
	return config
}
