package grpc

import (
	"context"
	"fmt"

	"github.com/kotche/gophKeeper/internal/pb"
)

// Login registers a new user
func (s *Sender) Login(username, password string) error {
	if username == "" || password == "" {
		return fmt.Errorf("login or password is empty")
	}

	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return fmt.Errorf("server is not available: %s", err.Error())
	}
	defer conn.Close()
	c := pb.NewAuthServiceClient(conn)
	r := &pb.UserRequest{Username: username, Password: password}

	ctx := context.Background()
	resp, err := c.Login(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("reg user id: %d, token: %s", resp.Id, resp.Token)
	if err = s.Service.SetUserParams(int(resp.Id), resp.Token); err != nil {
		return err
	}
	return nil
}

// Authentication identifies the user
func (s *Sender) Authentication(login, password string) error {
	if login == "" || password == "" {
		return fmt.Errorf("login or password is empty")
	}

	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return fmt.Errorf("server is not available: %s", err.Error())
	}
	defer conn.Close()
	c := pb.NewAuthServiceClient(conn)
	r := &pb.UserRequest{Username: login, Password: password}

	ctx := context.Background()
	resp, err := c.Authentication(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("auth user id: %d, token: %s", resp.Id, resp.Token)
	if err = s.Service.SetUserParams(int(resp.Id), resp.Token); err != nil {
		return err
	}
	return nil
}
