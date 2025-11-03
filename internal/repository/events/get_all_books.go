package events

import (
	"context"
	"fmt"

	"github.com/avraam311/event-booker/internal/models"
)

func (r *Repository) GetAllBooks(ctx context.Context) ([]*models.BookDB, error) {
	query := `
		SELECT id, created_at, event_id
		FROM book
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository/get_all_books.go - failed to get all books - %w", err)
	}
	defer rows.Close()

	var books []*models.BookDB
	for rows.Next() {
		var b models.BookDB
		err := rows.Scan(&b.ID, &b.CreatedAt, &b.EventID)
		if err != nil {
			return nil, fmt.Errorf("repository/get_all_books.go - failed to scan book row - %w", err)
		}
		books = append(books, &b)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository/get_all_books.go - failed to iterate books rows - %w", err)
	}

	return books, nil
}
