package events

import (
	"context"

	"github.com/avraam311/event-booker/internal/models"

	"github.com/go-playground/validator/v10"
)

type Service interface{
	CreateEvent(context.Context, *models.EventDTO) (uint, error)
	BookSeat(context.Context, uint, *models.BookDTO) (uint, error)
	Confirm(context.Context, uint) error
	GetEventInfo(context.Context, uint) (*models.EventDB, error)
}

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}
