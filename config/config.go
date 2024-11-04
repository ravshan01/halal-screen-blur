package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"os"
)

var cfg *Config = &Config{}

type Config struct {
	Port                int `env:"PORT" envDefault:"5000"`
	MaxImagesPerRequest int `env:"MAX_IMAGES_PER_REQUEST" envDefault:"10"`
	MaxImageSize        int `env:"MAX_IMAGE_SIZE" envDefault:"1000000"`
}

func GetConfig() *Config {
	return cfg
}

func LoadConfig() (*Config, error) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "dev"
	}

	err := godotenv.Load(".env." + appEnv)
	if err != nil {
		return nil, err
	}

	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
