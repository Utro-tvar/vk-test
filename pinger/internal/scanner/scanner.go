package scanner

import (
	"fmt"
	"log/slog"
	"net"

	docker "github.com/fsouza/go-dockerclient"
)

type Scanner struct {
	client *docker.Client
	logger *slog.Logger
}

func New(logger *slog.Logger) (*Scanner, error) {
	const op = "scanner.New"

	client, err := docker.NewClientFromEnv()
	if err != nil {
		return nil, fmt.Errorf("%s: Error while creating docker api client: %w", op, err)
	}
	return &Scanner{client: client, logger: logger}, nil
}

func (s *Scanner) Scan() []net.IP {
	const op = "scanner.Scan"

	containers, err := s.client.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s: Error while fetching containers", op), slog.Any("error", err))
		return nil
	}

	ips := make([]net.IP, 0, len(containers))

	for _, cont := range containers {
		for netName, network := range cont.Networks.Networks {
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
