package service

import (
	"sort"

	"github.com/kotche/gophKeeper/internal/client/domain"
)

func (s *Service) AddLoginPassword(data *domain.LoginPass) error {
	return s.Storage.AddLoginPassword(data)
}

func (s *Service) UpdateLoginPassword(data *domain.LoginPass) error {
	return s.Storage.UpdateLoginPassword(data)
}

func (s *Service) DeleteLoginPassword(id int) error {
	return s.Storage.DeleteLoginPassword(id)
}

func (s *Service) ReadAllLoginPasswordCache() ([]*domain.LoginPass, error) {
	data, err := s.Storage.ReadAllLoginPassword()
	if err != nil {
		return data, err
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	return data, err
}

func (s *Service) UpdateAllLoginPassCache(data []*domain.LoginPass) error {
	return s.Storage.UpdateAllLoginPass(data)
}
