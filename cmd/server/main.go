package main

import (
	"github.com/kotche/gophKeeper/internal/server/app"
	"github.com/kotche/gophKeeper/internal/server/config"
	"github.com/kotche/gophKeeper/pkg/logger"
)

func main() {
	log := logger.Init()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("server configuration error")
		return
	}
	appKeeper := app.NewKeeper(cfg, log)
	appKeeper.Run()
}
