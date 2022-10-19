package grpc

import (
	"context"
	"net"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kotche/gophKeeper/config/server"
	domainClient "github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/service"
	"github.com/kotche/gophKeeper/internal/client/storage"
	mock_server "github.com/kotche/gophKeeper/internal/mocks/server"
	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain"
	grpcServer "github.com/kotche/gophKeeper/internal/server/server/grpc"
	service_server "github.com/kotche/gophKeeper/internal/server/service"
	grpcHandler "github.com/kotche/gophKeeper/internal/server/transport/grpc"
	"github.com/kotche/gophKeeper/logger"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type testParamsText struct {
	client pb.TextServiceClient
	conn   *grpc.ClientConn
	lis    *bufconn.Listener
}

func NewTestParamsText(ctx context.Context, t *testing.T) *testParamsText {
	bufSize := 1024 * 1024
	lis := bufconn.Listen(bufSize)

	testParams := &testParamsText{}
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(testParams.bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	client := pb.NewTextServiceClient(conn)

	testParams.conn = conn
	testParams.lis = lis
	testParams.client = client

	return testParams
}

func (t *testParamsText) initServer(cfgServer *server.Config, log *zerolog.Logger, authService service_server.IAuthService, dataService service_server.IData) {
	serverService := service_server.NewService(authService, dataService, nil)
	handler := grpcHandler.NewHandler(serverService, log, cfgServer)
	grpcSrv := grpcServer.NewServer(cfgServer, handler)

	go func() {
		if err := grpcSrv.Run(t.lis); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("gRCP test server run error")
		}
	}()
}

func (t *testParamsText) bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return t.lis.Dial()
}

func (t *testParamsText) GetCode(err error) codes.Code {
	status, ok := status.FromError(err)
	if ok {
		return status.Code()
	}
	return codes.OK
}

func TestSender_CreateText(t *testing.T) {
	userID := 10
	text := "dgdgdfgdfg435dfg345345"
	meta := "meta info"

	ctx := context.Background()

	testParams := NewTestParamsText(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.Text{
		UserID:   userID,
		Text:     text,
		MetaInfo: meta,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockITextRepo(control)
	repo.EXPECT().Create(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, repo, nil, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.TextRequest{UserId: int64(userID), Text: text, MetaInfo: meta}
	_, actualError := testParams.client.CreateText(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_UpdateText(t *testing.T) {
	userID := 10
	ItemID := 55
	text := "dgdgdfgdfg435345345"
	meta := "meta info"

	ctx := context.Background()

	testParams := NewTestParamsText(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.Text{
		ID:       ItemID,
		UserID:   userID,
		Text:     text,
		MetaInfo: meta,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockITextRepo(control)
	repo.EXPECT().Update(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, repo, nil, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.TextUpdateRequest{Id: int64(ItemID), UserId: int64(userID), Text: text, MetaInfo: meta}
	_, actualError := testParams.client.UpdateText(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_DeleteText(t *testing.T) {
	userID := 10
	ItemID := 55

	ctx := context.Background()

	testParams := NewTestParamsText(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.Text{
		ID:     ItemID,
		UserID: userID,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockITextRepo(control)
	repo.EXPECT().Delete(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, repo, nil, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.TextDeleteRequest{Id: int64(ItemID), UserId: int64(userID)}
	_, actualError := testParams.client.DeleteText(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_GetAllText(t *testing.T) {
	userID := 10

	ctx := context.Background()

	testParams := NewTestParamsText(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := []domain.Text{
		{ID: 1, UserID: userID, Text: "5555dfgdfgdfg", MetaInfo: "meta"},
		{ID: 2, UserID: userID, Text: "6666dfgdgdfgdfgdgdfg"},
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockITextRepo(control)
	repo.EXPECT().GetAll(gomock.Any(), userID).Return(data, nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, repo, nil, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.TextGetAllRequest{UserId: int64(userID)}
	resp, actualError := testParams.client.GetAllText(ctx, req)
	assert.Equal(t, nil, actualError)

	dataResp := make([]domain.Text, 0, len(resp.Texts))
	for _, v := range resp.Texts {
		dataResp = append(dataResp, domain.Text{
			ID:       int(v.Id),
			UserID:   userID,
			Text:     v.Text,
			MetaInfo: v.MetaInfo,
		})
	}

	assert.EqualValues(t, data, dataResp)
}

func TestSender_ReadTextCache(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := service.NewService(cache, nil, log)
	sender := NewSender(srvc, nil, nil, log)

	data := []*domainClient.Text{
		{ID: 1, Text: "555dfgdf", MetaInfo: "777"},
		{ID: 2, Text: "666dfgdf"},
		{ID: 3, Text: "888dfgdfgdfg"},
	}
	for _, v := range data {
		cache.Save(v)
	}
	dataResp, actualError := sender.ReadTextCache()
	assert.Equal(t, nil, actualError)
	assert.EqualValues(t, data, dataResp)
}
