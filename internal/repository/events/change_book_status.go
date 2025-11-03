package events

import "context"

const (
	BookConfirmed = "confirmed"
)

func (r *Repository) ChangeBookStatus(ctx context.Context, id uint) error {
	query := `
		UPDATE book
		SET payment = $1
		WHERE id = $2;
	`

	res, err := r.db.ExecContext(ctx, query, BookConfirmed, id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return ErrBookNotFound
	}

	return nil
}
