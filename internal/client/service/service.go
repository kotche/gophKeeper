package service

import (
	"errors"
	"fmt"
	"reflect"
	"sort"

	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/rs/zerolog"
)

type ICache interface {
	SetUserParams(userID int, token string) error
	GetCurrentUserID() (int, error)
	GetToken() (string, error)
	GetVersion() (int, error)
	SetVersion(version int) error
	IncVersion() error

	Save(data any) error
	Update(data any) error
	Delete(data any) error
	GetAll(dt dataType.DataType) (any, error)
	UpdateAll(data any) error
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

func (s *Service) Save(data any) error {
	return s.Storage.Save(data)
}

func (s *Service) Update(data any) error {
	return s.Storage.Update(data)
}

func (s *Service) Delete(data any) error {
	return s.Storage.Delete(data)
}

func (s *Service) GetAll(dt dataType.DataType) (any, error) {
	data, err := s.Storage.GetAll(dt)
	if err != nil {
		return nil, err
	}

	switch d := data.(type) {
	case []*domain.LoginPass:
		sort.Slice(d, func(i, j int) bool {
			return d[i].ID < d[j].ID
		})
		return d, nil
	case []*domain.Text:
		sort.Slice(d, func(i, j int) bool {
			return d[i].ID < d[j].ID
		})
		return d, nil
	case []*domain.Binary:
		sort.Slice(d, func(i, j int) bool {
			return d[i].ID < d[j].ID
		})
		return d, nil
	case []*domain.BankCard:
		sort.Slice(d, func(i, j int) bool {
			return d[i].ID < d[j].ID
		})
		return d, nil
	default:
		err = fmt.Errorf("unsupported type '%v'", reflect.TypeOf(data))
		s.Log.Err(err).Msg("service getAll error")
		return nil, err
	}
}

func (s *Service) UpdateAll(data any) error {
	return s.Storage.UpdateAll(data)
}
