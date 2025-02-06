package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type Sender struct {
	logger *slog.Logger
	host   string
}

type entry struct {
	IP   string `json:"ip"`
	Ping int    `json:"ping"`
}

func New(logger *slog.Logger, host string) *Sender {
	return &Sender{logger: logger, host: host}
}

func (s *Sender) Send(data map[string]int) {
	const op = "sender.Send"
	entries := make([]entry, 0, len(data))
	for ip, ping := range data {
		entries = append(entries, entry{IP: ip, Ping: ping})
	}

	jsonData, err := json.Marshal(entries)
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s: Cannot marshal data", op), slog.Any("error", err))
		return
	}

	url := fmt.Sprintf("http://%s/update", s.host)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s: Error while sending data to backend", op), slog.Any("error", err))
		return
	}

	switch resp.StatusCode {
	case http.StatusOK:
		s.logger.Info(fmt.Sprintf("%s: Statistics updatet successfuly.", op))
	case http.StatusBadRequest:
		s.logger.Warn(fmt.Sprintf("%s: Wrong request format", op), slog.Any("request", jsonData))
	}
	resp.Body.Close()
}
