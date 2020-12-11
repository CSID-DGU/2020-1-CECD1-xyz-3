package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host string `envconfig:"HOST" default:"0.0.0.0"`
	Port int    `envconfig:"PORT" default:"50051"`

	DatabaseURI     string        `envconfig:"DATABASE_URI"`
	DatabaseTimeout time.Duration `envconfig:"DATABASE_TIMEOUT" default:"10s"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	_ = godotenv.Load()
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
