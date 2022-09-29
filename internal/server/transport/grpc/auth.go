package grpc

import (
	"context"
	"errors"

	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Registration registers a new user
func (h *Handler) Registration(ctx context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {
	user := domain.User{
		Login:    r.Login,
		Password: r.Password,
	}

	err := h.Service.Auth.CreateUser(ctx, &user)
	if errors.As(err, &errs.ConflictLoginError{}) {
		return nil, status.Errorf(codes.AlreadyExists, "Registration : %s", err.Error())
	} else if err != nil {
		h.Log.Error().Err(err).Msg("Registration: CreateUser service error")
		return nil, status.Errorf(codes.AlreadyExists, "Registration error: %s", err.Error())
	}

	response := pb.UserResponse{Id: user.ID, Login: user.Login, Password: user.Password}
	return &response, nil
}
