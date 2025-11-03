package events

import (
	"context"
	"fmt"

	"github.com/avraam311/event-booker/internal/models"
)

func (r *Repository) CreateEvent(ctx context.Context, ev *models.EventDTO) (uint, error) {
	query := `
		INSERT INTO event (name, seat_number, seats_number_left)
		VALUES ($1, $2, $2)
		RETURNING id;
	`

	var id uint
	err := r.db.QueryRowContext(ctx, query, ev.Name, ev.SeatsNumber).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repository/create_event.go - failed to scan id")
	}

	return id, nil
}
