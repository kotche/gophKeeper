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
		h.Log.Info().Err(err).Msg("Registration error")
		return nil, status.Errorf(codes.AlreadyExists, "Registration : %s", err.Error())
	} else if err != nil {
		h.Log.Error().Err(err).Msg("Registration: CreateUser service error")
		return nil, status.Error(codes.Internal, "Internal error")
	}

	response := pb.UserResponse{Id: user.ID, Login: user.Login, Password: user.Password}
	return &response, nil
}

// Authentication identifies the user
func (h *Handler) Authentication(ctx context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {
	user := domain.User{
		Login:    r.Login,
		Password: r.Password,
	}
	err := h.Service.Auth.AuthenticationUser(ctx, &user)

	if errors.As(err, &errs.AuthenticationError{}) {
		return nil, status.Errorf(codes.Unauthenticated, "Authentication error: %s", err.Error())
	} else if err != nil {
		h.Log.Error().Err(err).Msg("Authentication service error")
		return nil, status.Errorf(codes.AlreadyExists, "Authentication error: %s", err.Error())
	}

	response := pb.UserResponse{Id: user.ID, Login: user.Login, Password: user.Password}
	return &response, nil
}
