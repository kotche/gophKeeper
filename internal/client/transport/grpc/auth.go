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
	r := &pb.UserRequest{Username: username, Password: password}

	ctx := context.Background()
	resp, err := s.ClientConn.Auth.Login(ctx, r)
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
	r := &pb.UserRequest{Username: login, Password: password}

	ctx := context.Background()
	resp, err := s.ClientConn.Auth.Authentication(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("auth user id: %d, token: %s", resp.Id, resp.Token)
	if err = s.Service.SetUserParams(int(resp.Id), resp.Token); err != nil {
		return err
	}
	return nil
}
