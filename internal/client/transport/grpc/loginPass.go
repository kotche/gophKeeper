package grpc

import (
	"context"
	"fmt"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/pb"
)

func (s *Sender) CreateLoginPass(login, password, meta string) error {
	if login == "" || password == "" {
		return fmt.Errorf("login or password is empty")
	}

	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return fmt.Errorf("server is not available: %s", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			s.Log.Err(err).Msg("createLoginPass conn close error")
		}
	}()

	c := pb.NewLoginPassServiceClient(conn)

	userID, err := s.Service.GetCurrentUserID()
	if err != nil {
		return err
	}
	r := &pb.LoginPassRequest{UserId: int64(userID), Username: login, Password: password, MetaInfo: meta}

	ctx := context.Background()
	resp, err := c.CreateLoginPass(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("type lp create, userID %d, id: %d", userID, resp.Id)

	lp := &domain.LoginPass{
		ID:       int(resp.Id),
		Login:    login,
		Password: password,
		MetaInfo: meta,
	}

	if err = s.Service.AddLoginPassword(lp); err != nil {
		s.Log.Err(err).Msgf("createLoginPass add to cache '%+v' error: %w", lp, err)
	}

	if err = s.Service.Storage.IncVersion(); err != nil {
		s.Log.Debug().Msgf("createLoginPass inc version error: %w", err)
	}

	return nil
}

func (s *Sender) ReadLoginPassCache() ([]*domain.LoginPass, error) {
	return s.Service.ReadAllLoginPasswordCache()
}

func (s *Sender) GetAllLoginPass(ctx context.Context) ([]*domain.LoginPass, error) {
	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return nil, fmt.Errorf("server is not available: %s", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			s.Log.Err(err).Msg("getAllLoginPass conn close error")
		}
	}()

	c := pb.NewLoginPassServiceClient(conn)

	userID, err := s.Service.GetCurrentUserID()
	if err != nil {
		return nil, err
	}
	r := &pb.GetAllRequest{UserId: int64(userID)}

	resp, err := c.GetAllLoginPass(ctx, r)
	if err != nil {
		return nil, err
	}

	lpPairs := make([]*domain.LoginPass, 0, len(resp.LoginPassPairs))
	for _, v := range resp.LoginPassPairs {
		lpPairs = append(lpPairs, &domain.LoginPass{
			ID:       int(v.Id),
			Login:    v.Login,
			Password: v.Password,
			MetaInfo: v.MetaInfo,
		})
	}

	return lpPairs, nil
}
