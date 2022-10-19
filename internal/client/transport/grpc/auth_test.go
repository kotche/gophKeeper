package grpc

import (
	"context"
	"net"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kotche/gophKeeper/config/server"
	mock_server "github.com/kotche/gophKeeper/internal/mocks/server"
	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/errs"
	grpcServer "github.com/kotche/gophKeeper/internal/server/server/grpc"
	service_server "github.com/kotche/gophKeeper/internal/server/service"
	grpcHandler "github.com/kotche/gophKeeper/internal/server/transport/grpc"
	"github.com/kotche/gophKeeper/logger"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type testParamsAuth struct {
	auth pb.AuthServiceClient
	conn *grpc.ClientConn
	lis  *bufconn.Listener
}

func NewTestParamsAuth(ctx context.Context, t *testing.T) *testParamsAuth {
	bufSize := 1024 * 1024
	lis := bufconn.Listen(bufSize)

	testParams := &testParamsAuth{}
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(testParams.bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	auth := pb.NewAuthServiceClient(conn)

	testParams.conn = conn
	testParams.lis = lis
	testParams.auth = auth

	return testParams
}

func (t *testParamsAuth) initServer(cfgServer *server.Config, log *zerolog.Logger, authService service_server.IAuthService) {
	serverService := service_server.NewService(authService, nil, nil)
	handler := grpcHandler.NewHandler(serverService, log, cfgServer)
	grpcSrv := grpcServer.NewServer(cfgServer, handler)

	go func() {
		if err := grpcSrv.Run(t.lis); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("gRCP test server run error")
		}
	}()
}

func (t *testParamsAuth) bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return t.lis.Dial()
}

func (t *testParamsAuth) GetCode(err error) codes.Code {
	status, ok := status.FromError(err)
	if ok {
		return status.Code()
	}
	return codes.OK
}

func TestSender_Login(t *testing.T) {
	username := "username"
	password := "password"

	ctx := context.Background()

	testParams := NewTestParamsAuth(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	user := domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().CreateUser(gomock.Any(), &user).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	authService := service_server.NewAuthService(authRepo, jwt, pe, log)
	testParams.initServer(cfgServer, log, authService)

	req := &pb.UserRequest{Username: username, Password: password}
	_, actualError := testParams.auth.Login(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_LoginFailed(t *testing.T) {
	username := "username"
	password := "password"

	ctx := context.Background()

	testParams := NewTestParamsAuth(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	user := domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().CreateUser(gomock.Any(), &user).Return(errs.ConflictLoginError{}).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	authService := service_server.NewAuthService(authRepo, jwt, pe, log)
	testParams.initServer(cfgServer, log, authService)

	req := &pb.UserRequest{Username: username, Password: password}
	_, actualError := testParams.auth.Login(ctx, req)
	actualCode := testParams.GetCode(actualError)
	assert.Equal(t, codes.AlreadyExists, actualCode)
}

func TestSender_Authentication(t *testing.T) {
	userID := 1
	username := "username"
	password := "password"

	ctx := context.Background()

	testParams := NewTestParamsAuth(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	user := domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().GetUserID(gomock.Any(), &user).Return(userID, nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	authService := service_server.NewAuthService(authRepo, jwt, pe, log)
	testParams.initServer(cfgServer, log, authService)

	req := &pb.UserRequest{Username: username, Password: password}
	resp, actualError := testParams.auth.Authentication(ctx, req)
	assert.Equal(t, nil, actualError)
	assert.Equal(t, int64(userID), resp.Id)
}

func TestSender_AuthenticationFailed(t *testing.T) {
	userID := 0
	username := "username"
	password := "password"

	ctx := context.Background()

	testParams := NewTestParamsAuth(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	user := domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().GetUserID(gomock.Any(), &user).Return(userID, errs.AuthenticationError{}).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	authService := service_server.NewAuthService(authRepo, jwt, pe, log)
	testParams.initServer(cfgServer, log, authService)

	req := &pb.UserRequest{Username: username, Password: password}
	_, actualError := testParams.auth.Authentication(ctx, req)
	actualCode := testParams.GetCode(actualError)
	assert.Equal(t, codes.Unauthenticated, actualCode)
}
