package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		TCP      `yaml:"tcp"`
		Postgres `yaml:"postgres"`
	}

	TCP struct {
		Port string `env-required:"true" yaml:"port" env:"TCP_PORT"`
	}

	Postgres struct {
		DSN string `env-required:"true" yaml:"dsn" env:"PG_DSN"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./internal/server/config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %s", err.Error())
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
