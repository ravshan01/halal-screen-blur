package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
)

var cfg *Config

type Config struct {
	Port string `env:"PORT" envDefault:"5000"`
}

func GetConfig() *Config {
	return cfg
}

func LoadConfig(filenames ...string) (*Config, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Println("Error loading env files:", err)
		return nil, err
	}

	err = env.Parse(&cfg)
	if err != nil {
		log.Println("Error parsing env:", err)
		return nil, err
	}

	return cfg, nil
}
