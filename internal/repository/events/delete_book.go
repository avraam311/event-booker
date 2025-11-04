package events

import (
	"context"
	"fmt"
)

func (r *Repository) DeleteBook(ctx context.Context, bookID uint, evID uint) error {
	tx, err := r.db.Master.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("repositore/delete_book.go - failed to begin transaction: %w", err)
	}

	deleteQuery := `
        DELETE
		FROM book
		WHERE id = $1;
    `
	res, err := tx.ExecContext(ctx, deleteQuery, bookID)
	if err != nil {
		tx.Rollback()
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
        WHERE id = $1 AND seats_number_left < seats_number;
    `
	res, err = tx.ExecContext(ctx, updateQuery, evID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("repository/delete_book.go - failed to update event seats left: %w", err)
	}
	rows, _ = res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("repository/delete_book.go - all seats are available or event not found")
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("repository/delete_book.go - failed to commit transaction: %w", err)
	}

	return nil
}
