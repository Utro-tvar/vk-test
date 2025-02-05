package scanner

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Scanner struct {
	client *client.Client
	logger *slog.Logger
}

func New(logger *slog.Logger) (*Scanner, error) {
	const op = "scanner.New"

	client, err := client.NewClientWithOpts()
	if err != nil {
		return nil, fmt.Errorf("%s: Error while creating docker api client: %w", op, err)
	}
	return &Scanner{client: client, logger: logger}, nil
}

func (s *Scanner) Scan() []net.IP {
	const op = "scanner.Scan"

	containers, err := s.client.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s: Error while fetching containers: %w", op, err))
		return nil
	}

	ips := make([]net.IP, 0, len(containers))

	for _, cont := range containers {
		for netName, network := range cont.NetworkSettings.Networks {
			ip := net.ParseIP(network.IPAddress)
			if ip == nil {
				s.logger.Error(fmt.Sprintf("%s: Cannot parse ip: %s (network: %s, container: %s)", op, network.IPAddress, netName, cont.ID))
				continue
			}

			ips = append(ips, ip)
		}
	}

	return ips
}
