package service

import (
	"errors"

	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/rs/zerolog"
)

type ICache interface {
	SetUserParams(userID int, token string) error
	GetCurrentUserID() (int, error)
	GetToken() (string, error)
	GetVersion() (int, error)
	SetVersion(version int) error
	IncVersion() error
	AddLoginPassword(data *domain.LoginPass) error
	ReadAllLoginPassword() ([]*domain.LoginPass, error)
	UpdateAllLoginPass(lpPairs []*domain.LoginPass) error
}

type Service struct {
	Storage ICache
	Conf    *client.Config
	Log     *zerolog.Logger
}

func NewService(storage ICache, conf *client.Config, log *zerolog.Logger) *Service {
	return &Service{
		Storage: storage,
		Conf:    conf,
		Log:     log,
	}
}

func (s *Service) SetUserParams(userID int, token string) error {
	if userID < 0 {
		s.Log.Debug().Msgf("service setUserParams: userID %d less zero", userID)
		return errors.New("internal error")
	}

	if token == "" {
		s.Log.Debug().Msg("service setUserParams: token is empty")
		return errors.New("internal error")
	}

	if err := s.Storage.SetUserParams(userID, token); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetCurrentUserID() (int, error) {
	return s.Storage.GetCurrentUserID()
}

func (s *Service) GetToken() (string, error) {
	return s.Storage.GetToken()
}

func (s *Service) GetVersionCache() (int, error) {
	return s.Storage.GetVersion()
}

func (s *Service) SetVersionCache(version int) error {
	return s.Storage.SetVersion(version)
}
