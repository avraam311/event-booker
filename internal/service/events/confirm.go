package events

import (
	"context"
	"fmt"
)

func (s *Service) Confirm(ctx context.Context, id uint) error {
	err := s.repo.ChangeBookStatus(ctx, id)
	if err != nil {
		return fmt.Errorf("service/confirm.go - %w", err)
	}

	return nil
}
