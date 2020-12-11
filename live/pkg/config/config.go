package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Host      string `envconfig:"HOST" default:"0.0.0.0"`
	Port      int    `envconfig:"PORT" default:"8080"`
	CacheHost string `envconfig:"CACHE_HOST" default:"0.0.0.0"`
	CachePort int    `envconfig:"CACHE_PORT" default:"6379"`
}

func New() (Config, error) {
	var c Config
	return c, envconfig.Process("", &c)
}
