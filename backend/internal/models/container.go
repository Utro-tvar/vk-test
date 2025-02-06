package models

import (
	"net"
	"time"
)

type Container struct {
	IP             net.IP    `json:"ip"`
	Ping           int       `json:"ping"`
	LastConnection time.Time `json:"last_conn"`
}
