package client

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		GRPCClient `yaml:"grpc_client"`
		Logger     `yaml:"logger"`
		Updater    `yaml:"updater"`
	}

	GRPCClient struct {
		Address string        `env-required:"true" yaml:"grpc_address" env:"GRPC_ADDRESS"`
		Time    time.Duration `env-required:"true" yaml:"grpc_time" env:"GRPC_TIME"`
		Timeout time.Duration `env-required:"true" yaml:"grpc_timeout" env:"GRPC_TIMEOUT"`
	}

	Logger struct {
		LogLevel string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	Updater struct {
		Timeout time.Duration `env-required:"true" yaml:"update_timeout" env:"UPDATE_TIMEOUT"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	var configFilePath string
	if flag.Lookup("c") == nil {
		flag.StringVar(&configFilePath, "c", configFilePath, "config client file")
	}
	flag.Parse()

	if configFilePath == "" {
		return nil, errors.New("path config file is empty")
	}

	err := cleanenv.ReadConfig(configFilePath, cfg)
	if err != nil {
		return nil, fmt.Errorf("client error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
