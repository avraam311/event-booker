package events

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/avraam311/event-booker/internal/models"
)

func (r *Repository) GetEvent(ctx context.Context, id uint) (*models.EventDB, error) {
	query := `
        SELECT id, name, seats_number_left
        FROM event
        WHERE id = $1;
    `

	var ev models.EventDB
	err := r.db.QueryRowContext(ctx, query, id).Scan(&ev.ID, &ev.Name, &ev.SeatsNumberLeft)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEventNotFound
		}
		return nil, fmt.Errorf("repository/get_event.go - failed to get event - %w", err)
	}

	return &ev, nil
}
