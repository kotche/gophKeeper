package service

import (
	"github.com/kotche/gophKeeper/internal/client/config"
	"github.com/rs/zerolog"
)

type CacheContract interface {
}

type Service struct {
	storage CacheContract
	Conf    *config.Config
	Log     *zerolog.Logger
}

func NewService(conf *config.Config, log *zerolog.Logger) *Service {
	return &Service{Conf: conf, Log: log}
}
