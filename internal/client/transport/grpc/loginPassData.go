package grpc

import (
	"context"
	"fmt"

	"github.com/kotche/gophKeeper/internal/pb"
)

func (s *Sender) CreateLoginPass(username, password, meta string) error {
	if username == "" || password == "" {
		return fmt.Errorf("login or password is empty")
	}

	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log)
	if err != nil {
		return fmt.Errorf("server is not available: %s", err.Error())
	}
	defer conn.Close()
	c := pb.NewLoginPassServiceClient(conn)

	userID, err := s.Service.GetCurrentUserID()
	if err != nil {
		return err
	}
	r := &pb.LoginPassRequest{UserId: int64(userID), Username: username, Password: password, MetaInfo: meta}

	ctx := context.Background()
	resp, err := c.CreateLoginPass(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("type lp create, userID %d, id: %d", userID, resp.Id)
	return nil
}
