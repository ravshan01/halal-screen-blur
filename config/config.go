package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var cfg *Config = &Config{}

type Config struct {
	Port int `env:"PORT" envDefault:"5000"`
}

func GetConfig() *Config {
	return cfg
}

func LoadConfig(filenames ...string) (*Config, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		return nil, err
	}

	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
