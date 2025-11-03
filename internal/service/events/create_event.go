package events

import (
	"context"
	"fmt"

	"github.com/avraam311/event-booker/internal/models"
)

func (s *Service) CreateEvent(ctx context.Context, ev *models.EventDTO) (uint, error) {
	id, err := s.repo.CreateEvent(ctx, ev)
	if err != nil {
		return 0, fmt.Errorf("service/create_event.go - %w", err)
	}

	return id, nil
}
