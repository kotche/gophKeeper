package service

import (
	"sort"

	"github.com/kotche/gophKeeper/internal/client/domain"
)

func (s *Service) AddText(data *domain.Text) error {
	return s.Storage.AddText(data)
}

func (s *Service) UpdateText(data *domain.Text) error {
	return s.Storage.UpdateText(data)
}

func (s *Service) DeleteText(id int) error {
	return s.Storage.DeleteText(id)
}

func (s *Service) ReadAllTextCache() ([]*domain.Text, error) {
	data, err := s.Storage.ReadAllText()
	if err != nil {
		return data, err
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	return data, err
}

func (s *Service) UpdateAllTextCache(data []*domain.Text) error {
	return s.Storage.UpdateAllText(data)
}
