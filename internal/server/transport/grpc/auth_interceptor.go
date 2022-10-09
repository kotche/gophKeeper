package grpc

import (
	"context"
	"strings"

	"github.com/kotche/gophKeeper/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryAuthorize returns a server interceptor function authorize unary RPC
func (h *Handler) UnaryAuthorize(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//h.Log.Debug().Msgf("UnaryAuthorize interceptor: %s", info.FullMethod)

	if strings.Contains(info.FullMethod, "/Login") || strings.Contains(info.FullMethod, "/Authentication") {
		return handler(ctx, req)
	}

	userID, err := h.getUserIDFromRequest(req)
	if err != nil {
		return nil, err
	}
	err = h.authorize(ctx, userID)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func (h *Handler) authorize(ctx context.Context, userID int) error {
	const fInfo = "handler server authorize"

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		h.Log.Debug().Msgf("%s metadata is not provided", fInfo)
		return status.Errorf(codes.Unauthenticated, "authorization error")
	}

	values := md.Get("token")
	if len(values) == 0 {
		h.Log.Debug().Msgf("%s authorization token is not provided", fInfo)
		return status.Errorf(codes.Unauthenticated, "authorization error")
	}

	accessToken := values[0]
	claims, err := h.Service.Auth.Verify(accessToken)
	if err != nil {
		h.Log.Err(err).Msgf("%s access token is invalid", fInfo)
		return status.Error(codes.Unauthenticated, "authorization error")
	}

	if claims.ID != userID {
		h.Log.Debug().Msgf("%s user id claims '%d' not equal user id req '%d'", fInfo, claims.ID, userID)
		return status.Error(codes.Unauthenticated, "authorization error")
	}

	return nil
}

func (h *Handler) getUserIDFromRequest(req interface{}) (int, error) {
	var userID int64

	switch req.(type) {
	case *pb.LoginPassRequest:
		userID = req.(*pb.LoginPassRequest).UserId
	case *pb.GetAllRequest:
		userID = req.(*pb.GetAllRequest).UserId
	default:
		h.Log.Debug().Msg("handler server getUserIDFromRequest request unsupported type for get user id")
		return -1, status.Errorf(codes.Internal, "internal error")
	}

	return int(userID), nil
}
