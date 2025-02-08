package config

import (
	"fmt"
	"net"
	"os"
	"strconv"

	docker "github.com/fsouza/go-dockerclient"
)

type Config struct {
	PingPeriod  int    //in seconds
	BackAddress net.IP //address of backend service
	BackPort    int
}

func FromEnv() (*Config, error) {
	config := Config{}

	period, exist := os.LookupEnv("PING_PERIOD")
	if !exist {
		return nil, fmt.Errorf("no environment variable PING_PERIOD")
	}
	var err error
	config.PingPeriod, err = strconv.Atoi(period)
	if err != nil {
		return nil, fmt.Errorf("wrong format of environment variable PING_PERIOD")
	}

	backName, exist := os.LookupEnv("BACK_ADDR")
	if !exist {
		return nil, fmt.Errorf("no environment variable BACK_ADDR")
	}
	config.BackAddress, err = getIP(backName)
	if err != nil {
		return nil, fmt.Errorf("problems with searching IP of container %s", backName)
	}

	port, exist := os.LookupEnv("BACK_PORT")
	if !exist {
		return nil, fmt.Errorf("no environment variable BACK_PORT")
	}
	config.BackPort, err = strconv.Atoi(port)
	if config.BackPort < 0 || err != nil {
		return nil, fmt.Errorf("incorrect value of BACK_PORT")
	}

	return &config, nil
}

func getIP(name string) (net.IP, error) {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return nil, err
	}

	containers, err := client.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		return nil, fmt.Errorf("error while fetching containers: %w", err)
	}

	for _, cont := range containers {
		if cont.Names[0][1:] != name {
			continue
		}
		for _, network := range cont.Networks.Networks {
			ip := net.ParseIP(network.IPAddress)
			if ip == nil {
				return nil, fmt.Errorf("cannot parse ip of container with name %s", name)
			}
			return ip, nil
		}
	}

	return nil, fmt.Errorf("cannot find container with name %s", name)
}
