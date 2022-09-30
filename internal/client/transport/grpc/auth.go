package grpc

import (
	"context"
	"fmt"

	"github.com/kotche/gophKeeper/internal/pb"
)

// Registration registers a new user
func (s *Sender) Registration(login, password string) error {
	if login == "" || password == "" {
		return fmt.Errorf("login or password is empty")
	}

	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log)
	if err != nil {
		return fmt.Errorf("server is not available: %s", err.Error())
	}
	defer conn.Close()
	c := pb.NewAuthServiceClient(conn)
	r := &pb.UserRequest{Login: login, Password: password}

	ctx := context.Background()
	resp, err := c.Registration(ctx, r)
	if err != nil {
		return err
	}

	//TODO: implement writing to the storage on the client
	_ = resp

	return nil
}

// Authentication identifies the user
func (s *Sender) Authentication(login, password string) error {
	if login == "" || password == "" {
		return fmt.Errorf("login or password is empty")
	}

	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log)
	if err != nil {
		return fmt.Errorf("server is not available: %s", err.Error())
	}
	defer conn.Close()
	c := pb.NewAuthServiceClient(conn)
	r := &pb.UserRequest{Login: login, Password: password}

	ctx := context.Background()
	resp, err := c.Authentication(ctx, r)
	if err != nil {
		return err
	}

	//TODO: implement writing to the storage on the client
	_ = resp

	return nil
}
