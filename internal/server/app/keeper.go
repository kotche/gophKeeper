package app

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kotche/gophKeeper/config/server"
	grpcServer "github.com/kotche/gophKeeper/internal/server/server/grpc"
	"github.com/kotche/gophKeeper/internal/server/service"
	"github.com/kotche/gophKeeper/internal/server/storage"
	"github.com/kotche/gophKeeper/internal/server/storage/postgres"
	grpcHandler "github.com/kotche/gophKeeper/internal/server/transport/grpc"
	"github.com/rs/zerolog"
)

// Keeper application
type Keeper struct {
	Log *zerolog.Logger
	Cfg *server.Config
}

func NewKeeper(cfg *server.Config, log *zerolog.Logger) *Keeper {
	return &Keeper{Log: log, Cfg: cfg}
}

// Run start application, initializing dependencies
func (k *Keeper) Run() {
	pgx, err := postgres.NewPGX(k.Cfg.DSN)
	if err != nil {
		k.Log.Fatal().Err(err).Msg("db connection error")
	}
	commonRepo := postgres.NewCommonPostgres(pgx.DB, k.Log)
	authRepo := postgres.NewAuthPostgres(pgx.DB, k.Log)
	lpRepo := postgres.NewLoginPassPostgres(pgx.DB, k.Log)
	textRepo := postgres.NewTextPostgres(pgx.DB, k.Log)
	binaryRepo := postgres.NewBinaryPostgres(pgx.DB, k.Log)
	bankCardRepo := postgres.NewBankCardPostgres(pgx.DB, k.Log)
	repo := storage.NewRepository(commonRepo, authRepo, lpRepo, textRepo, binaryRepo, bankCardRepo)

	jwt := service.NewJWTManager(k.Cfg.SecretKeyToken, k.Cfg.TokenDuration, k.Log)
	authService := service.NewAuthService(repo.Auth, k.Log, jwt, k.Cfg.SecretKeyPassword)
	commonService := service.NewCommonService(commonRepo, k.Log)
	dataService := service.NewDataService(repo.LoginPass, repo.Text, repo.Binary, repo.BankCard, k.Log)
	srvc := service.NewService(commonService, authService, dataService)

	handler := grpcHandler.NewHandler(srvc, k.Log, k.Cfg)
	grpcSrv := grpcServer.NewServer(k.Cfg, handler)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		k.Log.Info().Msg("Starting gRPC server")
		if err = grpcSrv.Run(); err != nil && err != http.ErrServerClosed {
			k.Log.Fatal().Err(err).Msg("gRCP server run error")
		}
	}()
	<-termChan

	grpcSrv.Stop()
	k.Log.Info().Msg("Shutdown gRPC server")
}
