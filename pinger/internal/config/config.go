package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PingPeriod  int    //in seconds
	BackAddress string //address of backend service
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

	config.BackAddress, exist = os.LookupEnv("BACK_ADDR")
	if !exist {
		return nil, fmt.Errorf("no environment variable BACK_ADDR")
	}

	return &config, nil
}
