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

const handlerAuthorize = "handler server authorize"

// UnaryAuthorize gets a server interceptor function authorize unary RPC
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

// authorize checks the token for validity by user_id
func (h *Handler) authorize(ctx context.Context, userID int) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		h.Log.Debug().Msgf("%s metadata is not provided", handlerAuthorize)
		return status.Errorf(codes.Unauthenticated, "authorization error")
	}

	values := md.Get("token")
	if len(values) == 0 {
		h.Log.Debug().Msgf("%s authorization token is not provided", handlerAuthorize)
		return status.Errorf(codes.Unauthenticated, "authorization error")
	}

	accessToken := values[0]
	claims, err := h.Service.Auth.Verify(accessToken)
	if err != nil {
		h.Log.Err(err).Msgf("%s access token is invalid", handlerAuthorize)
		return status.Error(codes.Unauthenticated, "authorization error")
	}

	if claims.ID != userID {
		h.Log.Debug().Msgf("%s user id claims '%d' not equal user id req '%d'", handlerAuthorize, claims.ID, userID)
		return status.Error(codes.Unauthenticated, "authorization error")
	}

	return nil
}

func (h *Handler) getUserIDFromRequest(req interface{}) (int, error) {
	var userID int64

	switch req.(type) {
	case *pb.GetVersionRequest:
		userID = req.(*pb.GetVersionRequest).UserId

	case *pb.LoginPassRequest:
		userID = req.(*pb.LoginPassRequest).UserId
	case *pb.LoginPassUpdateRequest:
		userID = req.(*pb.LoginPassUpdateRequest).UserId
	case *pb.LoginPassDeleteRequest:
		userID = req.(*pb.LoginPassDeleteRequest).UserId
	case *pb.LoginPassGetAllRequest:
		userID = req.(*pb.LoginPassGetAllRequest).UserId

	case *pb.TextRequest:
		userID = req.(*pb.TextRequest).UserId
	case *pb.TextUpdateRequest:
		userID = req.(*pb.TextUpdateRequest).UserId
	case *pb.TextDeleteRequest:
		userID = req.(*pb.TextDeleteRequest).UserId
	case *pb.TextGetAllRequest:
		userID = req.(*pb.TextGetAllRequest).UserId

	case *pb.BinaryRequest:
		userID = req.(*pb.BinaryRequest).UserId
	case *pb.BinaryUpdateRequest:
		userID = req.(*pb.BinaryUpdateRequest).UserId
	case *pb.BinaryDeleteRequest:
		userID = req.(*pb.BinaryDeleteRequest).UserId
	case *pb.BinaryGetAllRequest:
		userID = req.(*pb.BinaryGetAllRequest).UserId

	case *pb.BankCardRequest:
		userID = req.(*pb.BankCardRequest).UserId
	case *pb.BankCardUpdateRequest:
		userID = req.(*pb.BankCardUpdateRequest).UserId
	case *pb.BankCardDeleteRequest:
		userID = req.(*pb.BankCardDeleteRequest).UserId
	case *pb.BankCardGetAllRequest:
		userID = req.(*pb.BankCardGetAllRequest).UserId

	default:
		h.Log.Debug().Msg("handler server getUserIDFromRequest request unsupported type for get user id")
		return -1, status.Errorf(codes.Internal, "internal error")
	}

	return int(userID), nil
}
