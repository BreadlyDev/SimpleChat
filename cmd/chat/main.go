package main

import (
	"log/slog"
	"os"
	"simplechat/internal/config"
	"simplechat/internal/lib/logger/sl"
	"simplechat/internal/storage/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

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

	// TODO: Add Migrator & Migrations

	// TODO: Initialize Server

	// TODO: Initialize App

	// TODO: ...
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
