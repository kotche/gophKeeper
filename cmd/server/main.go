package main

import (
	"github.com/kotche/gophKeeper/config/server"
	"github.com/kotche/gophKeeper/internal/server/app"
	"github.com/kotche/gophKeeper/logger"
)

func main() {
	log := logger.Init()
	cfg, err := server.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("server configuration error")
		return
	}
	appKeeper := app.NewKeeper(cfg, log)
	appKeeper.Run()
}
