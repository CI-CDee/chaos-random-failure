package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Set default values if not present
	setDefault("CHAOS_LATENCY_MS", "100")
	setDefault("CHAOS_ERROR_RATE", "0.2")
	setDefault("CHAOS_CRASH_RATE", "0.05")
	setDefault("PORT", "8080")
}

// setDefault ensures env variable is set
func setDefault(key, value string) {
	if os.Getenv(key) == "" {
		os.Setenv(key, value)
	}
}
