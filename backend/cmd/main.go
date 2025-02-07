package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Utro-tvar/vk-test/backend/internal/config"
	"github.com/Utro-tvar/vk-test/backend/internal/service"
	"github.com/Utro-tvar/vk-test/backend/internal/storage/postgres"
	"github.com/Utro-tvar/vk-test/backend/internal/transport/rest"
)

func main() {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)

	cfg, err := config.FromEnv()
	if err != nil {
		logger.Error("Cannot read config", slog.Any("error", err))
		os.Exit(1)
	}

	strg, err := postgres.New(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	if err != nil {
		logger.Error("Cannot initialize postgres", slog.Any("error", err))
		os.Exit(1)
	}

	serv := service.New(logger, strg)

	rest := rest.New(logger, serv)

	go rest.MustRun(":80")

	logger.Info("Service start listening at localhost:80")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done

	strg.Close()

	logger.Info("Server stopped")
}
