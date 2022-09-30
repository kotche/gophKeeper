package client

import (
	"errors"
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		TCP `yaml:"tcp"`
	}

	TCP struct {
		Port string `env-required:"true" yaml:"port" env:"TCP_PORT"`
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
