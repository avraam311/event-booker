package events

import (
	"context"
	"fmt"

	"github.com/avraam311/event-booker/internal/models"
)

func (s *Service) BookSeat(ctx context.Context, id uint, book *models.BookDTO) (uint, error) {
	id, err := s.repo.CreateBook(ctx, id, book)
	if err != nil {
		return 0, fmt.Errorf("service/book_seat.go - %w", err)
	}

	return id, nil
}
