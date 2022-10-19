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

type testParamsBankCard struct {
	client pb.BankCardServiceClient
	conn   *grpc.ClientConn
	lis    *bufconn.Listener
}

func NewTestParamsBankCard(ctx context.Context, t *testing.T) *testParamsBankCard {
	bufSize := 1024 * 1024
	lis := bufconn.Listen(bufSize)

	testParams := &testParamsBankCard{}
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(testParams.bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	client := pb.NewBankCardServiceClient(conn)

	testParams.conn = conn
	testParams.lis = lis
	testParams.client = client

	return testParams
}

func (t *testParamsBankCard) initServer(cfgServer *server.Config, log *zerolog.Logger, authService service_server.IAuthService, dataService service_server.IData) {
	serverService := service_server.NewService(authService, dataService, nil)
	handler := grpcHandler.NewHandler(serverService, log, cfgServer)
	grpcSrv := grpcServer.NewServer(cfgServer, handler)

	go func() {
		if err := grpcSrv.Run(t.lis); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("gRCP test server run error")
		}
	}()
}

func (t *testParamsBankCard) bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return t.lis.Dial()
}

func (t *testParamsBankCard) GetCode(err error) codes.Code {
	status, ok := status.FromError(err)
	if ok {
		return status.Code()
	}
	return codes.OK
}

func TestSender_CreateBankCard(t *testing.T) {
	userID := 10
	number := "555555"
	meta := "meta info"

	ctx := context.Background()

	testParams := NewTestParamsBankCard(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.BankCard{
		UserID:   userID,
		Number:   number,
		MetaInfo: meta,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBankCardRepo(control)
	repo.EXPECT().Create(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, nil, nil, repo, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.BankCardRequest{UserId: int64(userID), Number: number, MetaInfo: meta}
	_, actualError := testParams.client.CreateBankCard(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_UpdateBankCard(t *testing.T) {
	userID := 10
	ItemID := 55
	number := "555555"
	meta := "meta info"

	ctx := context.Background()

	testParams := NewTestParamsBankCard(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.BankCard{
		ID:       ItemID,
		UserID:   userID,
		Number:   number,
		MetaInfo: meta,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBankCardRepo(control)
	repo.EXPECT().Update(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, nil, nil, repo, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.BankCardUpdateRequest{Id: int64(ItemID), UserId: int64(userID), Number: number, MetaInfo: meta}
	_, actualError := testParams.client.UpdateBankCard(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_DeleteBankCard(t *testing.T) {
	userID := 10
	ItemID := 55

	ctx := context.Background()

	testParams := NewTestParamsBankCard(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.BankCard{
		ID:     ItemID,
		UserID: userID,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBankCardRepo(control)
	repo.EXPECT().Delete(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, nil, nil, repo, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.BankCardDeleteRequest{Id: int64(ItemID), UserId: int64(userID)}
	_, actualError := testParams.client.DeleteBankCard(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_GetAllBankCard(t *testing.T) {
	userID := 10

	ctx := context.Background()

	testParams := NewTestParamsBankCard(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := []domain.BankCard{
		{ID: 1, UserID: userID, Number: "5555", MetaInfo: "meta"},
		{ID: 2, UserID: userID, Number: "6666"},
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBankCardRepo(control)
	repo.EXPECT().GetAll(gomock.Any(), userID).Return(data, nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, nil, nil, repo, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.BankCardGetAllRequest{UserId: int64(userID)}
	resp, actualError := testParams.client.GetAllBankCard(ctx, req)
	assert.Equal(t, nil, actualError)

	dataResp := make([]domain.BankCard, 0, len(resp.BankCards))
	for _, v := range resp.BankCards {
		dataResp = append(dataResp, domain.BankCard{
			ID:       int(v.Id),
			UserID:   userID,
			Number:   v.Number,
			MetaInfo: v.MetaInfo,
		})
	}

	assert.EqualValues(t, data, dataResp)
}

func TestSender_ReadBankCardCache(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := service.NewService(cache, nil, log)
	sender := NewSender(srvc, nil, nil, log)

	data := []*domainClient.BankCard{
		{ID: 1, Number: "555", MetaInfo: "777"},
		{ID: 2, Number: "666"},
		{ID: 3, Number: "888"},
	}
	for _, v := range data {
		cache.Save(v)
	}
	dataResp, actualError := sender.ReadBankCardCache()
	assert.Equal(t, nil, actualError)
	assert.EqualValues(t, data, dataResp)
}
