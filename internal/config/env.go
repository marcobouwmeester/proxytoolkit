package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	ForwardAddr         string
	LogApiTiming        bool
	GenerateBrunoConfig bool
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		Port:                os.Getenv("PORT"),
		ForwardAddr:         os.Getenv("FORWARD_ADDR"),
		LogApiTiming:        os.Getenv("LOG_API_TIMING") == "true",
		GenerateBrunoConfig: os.Getenv("GENERATE_BRUNO_CONFIG") == "true",
	}

	if cfg.Port == "" || cfg.ForwardAddr == "" {
		log.Fatal("PORT and FORWARD_ADDR must be set")
	}

	return cfg
}
