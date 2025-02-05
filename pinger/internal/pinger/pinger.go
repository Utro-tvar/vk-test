package pinger

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

type Scanner interface {
	Scan() []net.IP
}

type Sender interface {
	Send(map[string]int)
}

type App struct {
	scanner Scanner
	sender  Sender
	logger  *slog.Logger
}

func pingLoop(a *App) {
	const op = "pinger.pingLoop"
	ips := a.scanner.Scan()

	errs := make(chan error, len(ips))
	wg := sync.WaitGroup{}
	wg.Add(len(ips))

	stats := make(map[string]int, len(ips))
	mx := sync.Mutex{}

	for _, ip := range ips {
		go func() {
			defer wg.Done()

			pinger, err := probing.NewPinger(ip.String())
			pinger.SetPrivileged(true)
			pinger.Timeout = 1 * time.Second

			if err != nil {
				errs <- fmt.Errorf("error while initializing of pinger on ip %s: %w", ip.String(), err)
				return
			}

			pinger.Count = 4
			err = pinger.Run()
			if err != nil {
				errs <- fmt.Errorf("error while pinging ip %s: %w", ip.String(), err)
				return
			}

			if pinger.Statistics().PacketsRecv > 0 {
				rtt := pinger.Statistics().AvgRtt.Milliseconds()

				mx.Lock()
				stats[ip.String()] = int(rtt)
				mx.Unlock()
			}
		}()
	}

	wg.Wait()

	close(errs)
	for err := range errs {
		if err != nil {
			a.logger.Error(op, slog.Any("error", err))
		}
	}

	a.sender.Send(stats)
}

func New(scanner Scanner, sender Sender, logger *slog.Logger) *App {
	return &App{scanner: scanner, sender: sender, logger: logger}
}

func (a *App) Run(ctx context.Context, pingPeriod int) {
	const op = "pinger.Run"
	a.logger.Info("Service started.")
	go func() {
		ticker := time.NewTicker(time.Duration(pingPeriod) * time.Second)
		defer ticker.Stop()

		pingLoop(a)
		for {
			select {
			case <-ctx.Done():
				a.logger.Info("Service stopped.")
				return
			case <-ticker.C:
				pingLoop(a)
			}
		}
	}()
}
