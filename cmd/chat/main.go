package main

import (
	"log/slog"
	"os"
	"simplechat/internal/config"
	"simplechat/internal/lib/logger/sl"
	"simplechat/internal/storage/postgres"
	"simplechat/internal/utils/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info(
		"starting simple-chat application",
		slog.String("env", cfg.Env),
	)

	storage, err := postgres.New(&cfg.Storage)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	err = storage.Ping()
	if err != nil {
		log.Error("failed to connect to database", sl.Err(err))
	} else {
		log.Info("successful connection")
	}

	_ = storage

	// TODO: Initialize Server

	// TODO: ...
}
