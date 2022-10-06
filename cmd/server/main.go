package main

import (
	"github.com/kotche/gophKeeper/config/server"
	"github.com/kotche/gophKeeper/internal/server/app"
	"github.com/kotche/gophKeeper/logger"
)

func main() {
	cfg, err := server.NewConfig()
	if err != nil {
		log := logger.Init("")
		log.Fatal().Err(err).Msg("server configuration error")
		return
	}
	log := logger.Init(cfg.LogLevel)
	appKeeper := app.NewKeeper(cfg, log)
	appKeeper.Run()
}
