package events

import (
	"context"
	"fmt"

	"github.com/avraam311/event-booker/internal/models"

	"github.com/lib/pq"
)

const (
	BookNotConfirmed = "not confirmed"
)

func (r *Repository) CreateBook(ctx context.Context, id uint, book *models.BookDTO) (uint, error) {
	query := `
		INSERT INTO book (person_name, event_id, book)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	var bookID uint
	err := r.db.QueryRowContext(ctx, query, book.PersonName, id, BookNotConfirmed).Scan(&bookID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			return 0, ErrEventNotFound
		}

		return 0, fmt.Errorf("repository/create_book.go - failed to scan id")
	}

	return bookID, nil
}
