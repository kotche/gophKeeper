package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/kotche/gophKeeper/internal/pb"
)

func (s *Sender) CreateLoginPass(login, password, meta string) (int, error) {
	userID := s.Service.GetCurrentUserID()
	r := &pb.LoginPassRequest{UserId: int64(userID), Username: login, Password: password, MetaInfo: meta}

	ctx := context.Background()
	resp, err := s.ClientConn.Lp.CreateLoginPass(ctx, r)
	if err != nil {
		return -1, err
	}

	s.Log.Debug().Msgf("type lp create, userID %d, id: %d", userID, resp.Id)

	data := &domain.LoginPass{
		ID:       int(resp.Id),
		Login:    login,
		Password: password,
		MetaInfo: meta,
	}

	if err = s.Service.Save(data); err != nil {
		s.Log.Err(err).Msgf("createLoginPass add to cache '%+v' error: %w", data, err)
	}

	s.Service.Storage.IncVersion()

	return data.ID, nil
}

func (s *Sender) UpdateLoginPass(id int, login, password, meta string) error {
	userID := s.Service.GetCurrentUserID()
	r := &pb.LoginPassUpdateRequest{Id: int64(id), UserId: int64(userID), Username: login, Password: password, MetaInfo: meta}

	ctx := context.Background()
	_, err := s.ClientConn.Lp.UpdateLoginPass(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("type lp update, userID %d, id: %d", userID, id)

	data := &domain.LoginPass{
		ID:       id,
		Login:    login,
		Password: password,
		MetaInfo: meta,
	}

	if err = s.Service.Update(data); err != nil {
		s.Log.Err(err).Msgf("updateLoginPass update lp to cache '%+v' error: %w", data, err)
	}

	s.Service.Storage.IncVersion()

	return nil
}

func (s *Sender) DeleteLoginPass(id int) error {
	userID := s.Service.GetCurrentUserID()
	r := &pb.LoginPassDeleteRequest{Id: int64(id), UserId: int64(userID)}

	ctx := context.Background()
	_, err := s.ClientConn.Lp.DeleteLoginPass(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("lp delete, userID %d, id: %d", userID, id)

	data := &domain.LoginPass{
		ID: id,
	}

	if err = s.Service.Delete(data); err != nil {
		s.Log.Err(err).Msgf("deleteLoginPass delete lp to cache '%d' error: %w", id, err)
	}

	s.Service.Storage.IncVersion()

	return nil
}

func (s *Sender) ReadLoginPassCache() ([]*domain.LoginPass, error) {
	data, err := s.Service.GetAll(dataType.LP)
	if err != nil {
		return nil, err
	}
	return data.([]*domain.LoginPass), nil
}

func (s *Sender) GetAllLoginPass(ctx context.Context) ([]*domain.LoginPass, error) {
	userID := s.Service.GetCurrentUserID()
	r := &pb.LoginPassGetAllRequest{UserId: int64(userID)}

	resp, err := s.ClientConn.Lp.GetAllLoginPass(ctx, r)
	if err != nil {
		return nil, err
	}

	data := make([]*domain.LoginPass, 0, len(resp.LoginPassPairs))
	for _, v := range resp.LoginPassPairs {
		data = append(data, &domain.LoginPass{
			ID:       int(v.Id),
			Login:    v.Login,
			Password: v.Password,
			MetaInfo: v.MetaInfo,
		})
	}

	return data, nil
}
