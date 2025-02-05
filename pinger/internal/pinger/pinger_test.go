package pinger_test

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"testing"
	"time"

	"github.com/Utro-tvar/vk-test/pinger/internal/pinger"
)

type TestSender struct {
	logger *slog.Logger
}

func (s *TestSender) Send(data map[string]int) {
	s.logger.Debug("Stats:")
	for k, v := range data {
		s.logger.Debug(fmt.Sprintf("%s: %d", k, v))
	}
}

type TestScanner struct{}

func (s *TestScanner) Scan() []net.IP {
	return []net.IP{net.ParseIP("12.123.48.5"), net.ParseIP("8.8.8.8")}
}

func TestPing(t *testing.T) {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)
	scanner := &TestScanner{}
	sender := &TestSender{logger: logger}

	app := pinger.New(scanner, sender, logger)

	ctx, cancel := context.WithCancel(context.Background())

	app.Run(ctx, 10)

	time.Sleep(time.Second * 16)

	cancel()

	time.Sleep(time.Second)
}
