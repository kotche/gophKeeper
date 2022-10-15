package server

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config configures server settings
	Config struct {
		GRPCServer `yaml:"grpc_server"`
		Postgres   `yaml:"postgres"`
		Security   `yaml:"security"`
		Logger     `yaml:"logger"`
	}

	GRPCServer struct {
		// Address TCP port server gRPC
		Address string `env-required:"true" yaml:"grpc_address" env:"GRPC_ADDRESS"`
	}

	Postgres struct {
		// DSN data source name postgres
		DSN string `env-required:"true" yaml:"dsn" env:"PG_DSN"`
	}

	Security struct {
		// TokenDuration token lifetime
		TokenDuration time.Duration `env-required:"true" yaml:"token_duration" env:"TOKEN_DURATION"`
		// SecretKeyToken secret key for token generation for user
		SecretKeyToken string `env-required:"true" yaml:"secret_key_token" env:"SECRET_KEY_TOKEN"`
		// SecretKeyPassword secret key for password generation for user
		SecretKeyPassword string `env-required:"true" yaml:"secret_key_password" env:"SECRET_KEY_PASSWORD"`
	}

	Logger struct {
		// LogLevel sets the logging level
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
