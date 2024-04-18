package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Clickhouse
	Router
}

type Clickhouse struct {
	Url string `env:"CLICKHOUSE_URL"`
}

type Router struct {
	ServerAddr string `env:"SERVER_ADDR"`
	ServerPort string `env:"SERVER_PORT"`
}

func GetConfig() *Config {
	cfg := Config{}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal().Msgf("error loading .env: %v", err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatal().Msgf("error parsing .env: %v", err)
	}

	return &cfg
}
