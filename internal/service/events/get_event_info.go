package events

import (
	"context"
	"fmt"

	"github.com/avraam311/event-booker/internal/models"
)

func (s *Service) GetEventInfo(ctx context.Context, id uint) (*models.EventDB, error) {
	ev, err := s.repo.GetEvent(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service/confirm.go - %w", err)
	}

	return ev, nil
}
