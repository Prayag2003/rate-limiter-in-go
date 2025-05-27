package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v10"
	"gopkg.in/yaml.v3"
)

type Config struct {
	RateLimiterType string `yaml:"rate_limiter_type" env:"RATE_LIMITER_TYPE" envDefault:"token"`
	Capacity        int    `yaml:"capacity" env:"CAPACITY" envDefault:"100"`
	RefillRate      int    `yaml:"refill_rate" env:"REFILL_RATE" envDefault:"50"`
	LeakRate        int    `yaml:"leak_rate" env:"LEAK_RATE" envDefault:"50"`
	RPS             int    `yaml:"rps" env:"RPS" envDefault:"100"`
	DurationSec     int    `yaml:"duration_sec" env:"DURATION_SEC" envDefault:"10"`
	Concurrency     int    `yaml:"concurrency" env:"CONCURRENCY" envDefault:"20"`
}

func LoadConfig() *Config {
	cfg := &Config{}

	if _, err := os.Stat("config.yaml"); err == nil {
		// YAML file exists
		f, err := os.Open("config.yaml")
		if err != nil {
			log.Fatalf("Error opening config.yaml: %v", err)
		}
		defer f.Close()

		decoder := yaml.NewDecoder(f)
		if err := decoder.Decode(cfg); err != nil {
			log.Fatalf("Error decoding config.yaml: %v", err)
		}
		log.Println("Loaded configuration from config.yaml")
	} else {
		// Fall back to ENV
		if err := env.Parse(cfg); err != nil {
			log.Fatalf("Failed to parse .env variables: %v", err)
		}
		log.Println("Loaded configuration from .env or env vars")
	}

	return cfg
}
