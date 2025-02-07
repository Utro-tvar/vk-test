package service

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/Utro-tvar/vk-test/backend/internal/models"
)

type Storage interface {
	Store(data []models.Container) error
	GetAll() ([]models.Container, error)
}

type Service struct {
	logger  *slog.Logger
	storage Storage
}

func New(logger *slog.Logger, storage Storage) *Service {
	return &Service{logger: logger, storage: storage}
}

func (s *Service) GetStatistics() ([]models.Container, error) {
	const op = "service.GetStatistics"

	cont, err := s.storage.GetAll()
	if err != nil {
		return cont, fmt.Errorf("%s: %w", op, err)
	} else {
		return cont, nil
	}
}

func (s *Service) UpdateStatistics(data []models.Container) {
	const op = "service.UpdateStatistics"

	for i := range data {
		data[i].LastConnection = models.JSONTime(time.Now())
	}

	err := s.storage.Store(data)
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s: error while storing stats", op), slog.Any("error", err))
	}
}
