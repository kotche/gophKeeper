package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/kotche/gophKeeper/internal/pb"
)

// CreateText creates text data
func (s *Sender) CreateText(text, meta string) (int, error) {
	userID := s.Service.GetCurrentUserID()
	r := &pb.TextRequest{UserId: int64(userID), Text: text, MetaInfo: meta}

	ctx := context.Background()
	resp, err := s.ClientConn.Text.CreateText(ctx, r)
	if err != nil {
		return -1, err
	}

	s.Log.Debug().Msgf("type text create, userID %d, id: %d", userID, resp.Id)

	data := &domain.Text{
		ID:       int(resp.Id),
		Text:     text,
		MetaInfo: meta,
	}

	if err = s.Service.Save(data); err != nil {
		s.Log.Err(err).Msgf("createText add to cache '%+v'", data)
	}

	s.Service.Storage.IncVersion()

	return data.ID, nil
}

// UpdateText updates text data
func (s *Sender) UpdateText(id int, text, meta string) error {
	userID := s.Service.GetCurrentUserID()
	r := &pb.TextUpdateRequest{Id: int64(id), UserId: int64(userID), Text: text, MetaInfo: meta}

	ctx := context.Background()
	_, err := s.ClientConn.Text.UpdateText(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("type lp update, userID %d, id: %d", userID, id)

	data := &domain.Text{
		ID:       id,
		Text:     text,
		MetaInfo: meta,
	}

	if err = s.Service.Update(data); err != nil {
		s.Log.Err(err).Msgf("updateText update text to cache '%+v'", data)
	}

	s.Service.Storage.IncVersion()

	return nil
}

// DeleteText deletes text data
func (s *Sender) DeleteText(id int) error {
	userID := s.Service.GetCurrentUserID()
	r := &pb.TextDeleteRequest{Id: int64(id), UserId: int64(userID)}

	ctx := context.Background()
	_, err := s.ClientConn.Text.DeleteText(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("lp delete, userID %d, id: %d", userID, id)

	data := &domain.Text{
		ID: id,
	}

	if err = s.Service.Delete(data); err != nil {
		s.Log.Err(err).Msgf("delete text to cache '%d'", id)
	}

	s.Service.Storage.IncVersion()

	return nil
}

// ReadTextCache reads text data from local repository
func (s *Sender) ReadTextCache() ([]*domain.Text, error) {
	data, err := s.Service.GetAll(dataType.TEXT)
	if err != nil {
		return nil, err
	}
	return data.([]*domain.Text), nil
}

// GetAllText gets text data from server db
func (s *Sender) GetAllText(ctx context.Context) ([]*domain.Text, error) {
	userID := s.Service.GetCurrentUserID()
	r := &pb.TextGetAllRequest{UserId: int64(userID)}

	resp, err := s.ClientConn.Text.GetAllText(ctx, r)
	if err != nil {
		return nil, err
	}

	data := make([]*domain.Text, 0, len(resp.Texts))
	for _, v := range resp.Texts {
		data = append(data, &domain.Text{
			ID:       int(v.Id),
			Text:     v.Text,
			MetaInfo: v.MetaInfo,
		})
	}

	return data, nil
}
