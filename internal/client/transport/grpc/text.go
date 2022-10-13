package grpc

import (
	"context"
	"fmt"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/kotche/gophKeeper/internal/pb"
)

func (s *Sender) CreateText(text, meta string) (int, error) {
	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return -1, fmt.Errorf("server is not available: %s", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			s.Log.Err(err).Msg("createText conn close error")
		}
	}()

	c := pb.NewTextServiceClient(conn)

	userID := s.Service.GetCurrentUserID()
	r := &pb.TextRequest{UserId: int64(userID), Text: text, MetaInfo: meta}

	ctx := context.Background()
	resp, err := c.CreateText(ctx, r)
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
		s.Log.Err(err).Msgf("createText add to cache '%+v' error: %w", data, err)
	}

	s.Service.Storage.IncVersion()

	return data.ID, nil
}

func (s *Sender) UpdateText(id int, text, meta string) error {
	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return fmt.Errorf("server is not available: %s", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			s.Log.Err(err).Msg("updateText conn close error")
		}
	}()

	c := pb.NewTextServiceClient(conn)

	userID := s.Service.GetCurrentUserID()
	r := &pb.TextUpdateRequest{Id: int64(id), UserId: int64(userID), Text: text, MetaInfo: meta}

	ctx := context.Background()
	_, err = c.UpdateText(ctx, r)
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
		s.Log.Err(err).Msgf("updateText update text to cache '%+v' error: %w", data, err)
	}

	s.Service.Storage.IncVersion()

	return nil
}

func (s *Sender) DeleteText(id int) error {
	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return fmt.Errorf("server is not available: %s", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			s.Log.Err(err).Msg("deleteText conn close error")
		}
	}()

	c := pb.NewTextServiceClient(conn)

	userID := s.Service.GetCurrentUserID()
	r := &pb.TextDeleteRequest{Id: int64(id), UserId: int64(userID)}

	ctx := context.Background()
	_, err = c.DeleteText(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("lp delete, userID %d, id: %d", userID, id)

	data := &domain.Text{
		ID: id,
	}

	if err = s.Service.Delete(data); err != nil {
		s.Log.Err(err).Msgf("delete text to cache '%d' error: %w", id, err)
	}

	s.Service.Storage.IncVersion()

	return nil
}

func (s *Sender) ReadTextCache() ([]*domain.Text, error) {
	data, err := s.Service.GetAll(dataType.TEXT)
	if err != nil {
		return nil, err
	}
	return data.([]*domain.Text), nil
}

func (s *Sender) GetAllText(ctx context.Context) ([]*domain.Text, error) {
	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return nil, fmt.Errorf("server is not available: %s", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			s.Log.Err(err).Msg("getAllText conn close error")
		}
	}()

	c := pb.NewTextServiceClient(conn)

	userID := s.Service.GetCurrentUserID()
	r := &pb.TextGetAllRequest{UserId: int64(userID)}

	resp, err := c.GetAllText(ctx, r)
	if err != nil {
		return nil, err
	}

	data := make([]*domain.Text, 0, len(resp.Text))
	for _, v := range resp.Text {
		data = append(data, &domain.Text{
			ID:       int(v.Id),
			Text:     v.Text,
			MetaInfo: v.MetaInfo,
		})
	}

	return data, nil
}
