package service

import (
	"fmt"
	"log/slog"

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

func (s *Service) GetStatistics() ([]models.Container, error) {
	const op = "service.GetStatistics"

	cont, err := s.storage.GetAll()
	return cont, fmt.Errorf("%s: %w", op, err)
}

func (s *Service) UpdateStatistics(data []models.Container) {
	const op = "service.UpdateStatistics"

	err := s.storage.Store(data)
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s: error while storing stats", op), slog.Any("error", err))
	}
}
