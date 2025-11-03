package events

import (
	"context"
	"fmt"

	"github.com/lib/pq"
)

func (r *Repository) DeleteBook(ctx context.Context, id uint) error {
	tx, err := r.db.Master.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("repositore/delete_book.go - failed to begin transaction: %w", err)
	}

	deleteQuery := `
        DELETE
		FROM book
		WHERE id = $1;
    `
	res, err := tx.ExecContext(ctx, deleteQuery, id)
	if err != nil {
		tx.Rollback()
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			return ErrBookNotFound
		}
		return fmt.Errorf("repository/delete_book.go - failed to delete book: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		return fmt.Errorf("repository/delete_book.go - no seats left or event not found")
	}

	updateQuery := `
        UPDATE event
        SET seats_number_left = seats_number_left + 1
        WHERE id = $1 AND seats_number_left < seat_number;
    `
	res, err = tx.ExecContext(ctx, updateQuery, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("repository/delete_book.go - failed to update event seats left: %w", err)
	}
	rows, _ = res.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		return fmt.Errorf("repository/delete_book.go - no seats left or event not found")
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("repository/delete_book.go - failed to commit transaction: %w", err)
	}

	return nil
}
