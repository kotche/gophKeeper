package app

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/service"
	"github.com/kotche/gophKeeper/internal/client/transport"
	grpcTransport "github.com/kotche/gophKeeper/internal/client/transport/grpc"
	"github.com/rs/zerolog"
)

type App struct {
	Conf *client.Config
	Log  *zerolog.Logger
}

func NewApp(conf *client.Config, log *zerolog.Logger) *App {
	return &App{Conf: conf, Log: log}
}

func (a *App) Run() {
	clientConn := grpcTransport.Connection{}
	srvc := service.NewService(a.Conf, a.Log)
	sender := grpcTransport.NewSender(srvc, clientConn, a.Conf, a.Log)
	commander := transport.NewCommander(sender, a.Conf, a.Log)

	fmt.Println("GophKeeper start")

	p := prompt.New(
		commander.Executor,
		commander.Completer,
		prompt.OptionTitle("menu"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Green),
	)
	p.Run()
}
