package models

import (
	"fmt"
	"net"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02"))
	return []byte(stamp), nil
}

type Container struct {
	IP             net.IP   `json:"ip"`
	Ping           int      `json:"ping"`
	LastConnection JSONTime `json:"last_conn"`
}
