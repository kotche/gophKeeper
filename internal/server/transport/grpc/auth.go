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

// Login registers a new user
func (h *Handler) Login(ctx context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {
	user := domain.User{
		Username: r.Username,
		Password: r.Password,
	}

	err := h.Service.Auth.CreateUser(ctx, &user)

	if errors.As(err, &errs.ConflictLoginError{}) {
		h.Log.Info().Err(err).Msg("handler login error")
		return nil, status.Errorf(codes.AlreadyExists, "login: %s", err.Error())
	} else if err != nil {
		h.Log.Error().Err(err).Msg("handler login error")
		return nil, status.Error(codes.Internal, "internal error")
	}

	response := pb.UserResponse{Id: int64(user.ID), Token: user.Token}
	return &response, nil
}

// Authentication identifies the user
func (h *Handler) Authentication(ctx context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {
	user := domain.User{
		Username: r.Username,
		Password: r.Password,
	}

	err := h.Service.Auth.AuthenticationUser(ctx, &user)

	if errors.As(err, &errs.AuthenticationError{}) {
		h.Log.Error().Err(err).Msg("handler authentication error")
		return nil, status.Errorf(codes.Unauthenticated, "authentication: %s", err.Error())
	} else if err != nil {
		h.Log.Error().Err(err).Msg("handler authentication error")
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	response := pb.UserResponse{Id: int64(user.ID), Token: user.Token}
	return &response, nil
}
