package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Utro-tvar/vk-test/pinger/internal/config"
	"github.com/Utro-tvar/vk-test/pinger/internal/pinger"
	"github.com/Utro-tvar/vk-test/pinger/internal/scanner"
	"github.com/Utro-tvar/vk-test/pinger/internal/sender"
)

func main() {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)

	logger.Info("Starting service")

	config, err := config.FromEnv()
	if err != nil {
		logger.Error("Reading config failed", slog.Any("error", err))
		os.Exit(1)
	}

	scanner, err := scanner.New(logger)
	if err != nil {
		logger.Error("Scanner initialization failed", slog.Any("error", err))
		os.Exit(1)
	}
	logger.Info("Scanner initialized successfuly")

	sender := sender.New(logger, fmt.Sprintf("%s:%d", config.BackAddress.String(), config.BackPort))
	logger.Info("Sender initialized successfuly")

	pinger := pinger.New(scanner, sender, logger)

	ctx, stop := context.WithCancel(context.Background())
	logger.Info("Running pinger...")
	pinger.Run(ctx, config.PingPeriod)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done

	logger.Info("Stopping service...")

	stop()
}
