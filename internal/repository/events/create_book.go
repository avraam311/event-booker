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
	tx, err := r.db.Master.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("repositore/create_book.go - failed to begin transaction: %w", err)
	}

	var bookID uint
	insertQuery := `
        INSERT INTO book (person_name, event_id, book)
        VALUES ($1, $2, $3)
        RETURNING id;
    `
	err = tx.QueryRowContext(ctx, insertQuery, book.PersonName, id, BookNotConfirmed).Scan(&bookID)
	if err != nil {
		tx.Rollback()
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			return 0, ErrEventNotFound
		}
		return 0, fmt.Errorf("repository/create_book.go - failed to create book: %w", err)
	}

	updateQuery := `
        UPDATE event
        SET seats_number_left = seats_number_left - 1
        WHERE id = $1 AND seats_number_left > 0;
    `
	res, err := tx.ExecContext(ctx, updateQuery, id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("repository/create_book.go - failed to update event seats left: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		return 0, ErrNoSeatsOrEventNotFound
	}
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("repository/create_book.go - failed to commit transaction: %w", err)
	}

	return bookID, nil
}
