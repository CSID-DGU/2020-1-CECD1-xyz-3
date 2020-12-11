package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host string `envconfig:"HOST" default:"0.0.0.0"`
	Port int    `envconfig:"PORT" default:"8080"`

	SecureServiceConnection    bool `envconfig:"SECURE_SERVICE_CONNECTION"`
	SecureGCPServiceConnection bool `envconfig:"SECURE_GCP_SERVICE_CONNECTION"`

	AuthServiceURI      string `envconfig:"AUTH_SERVICE_URI"`
	AuthServiceAudience string `envconfig:"AUTH_SERVICE_AUDIENCE"`
	JobServiceURI       string `envconfig:"JOB_SERVICE_URI"`
	JobServiceAudience  string `envconfig:"JOB_SERVICE_AUDIENCE"`
}

func New() (*Config, error) {
	cfg := Config{}
	_ = godotenv.Load()
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
