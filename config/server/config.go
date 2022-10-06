package server

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		TCP      `yaml:"tcp"`
		Postgres `yaml:"postgres"`
		Security `yaml:"security"`
		Logger   `yaml:"logger"`
	}

	TCP struct {
		Port string `env-required:"true" yaml:"port" env:"TCP_PORT"`
	}

	Postgres struct {
		DSN string `env-required:"true" yaml:"dsn" env:"PG_DSN"`
	}

	Security struct {
		TokenDuration     time.Duration `env-required:"true" yaml:"token_duration" env:"TOKEN_DURATION"`
		SecretKeyToken    string        `env-required:"true" yaml:"secret_key_token" env:"SECRET_KEY_TOKEN"`
		SecretKeyPassword string        `env-required:"true" yaml:"secret_key_password" env:"SECRET_KEY_PASSWORD"`
	}

	Logger struct {
		LogLevel string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	var configFilePath string
	if flag.Lookup("c") == nil {
		flag.StringVar(&configFilePath, "c", configFilePath, "config server file")
	}
	flag.Parse()

	if configFilePath == "" {
		return nil, errors.New("path config file is empty")
	}

	err := cleanenv.ReadConfig(configFilePath, cfg)
	if err != nil {
		return nil, fmt.Errorf("server error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
