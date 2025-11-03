package events

import (
	"context"

	"github.com/avraam311/event-booker/internal/models"
)

type Repository interface {
	CreateEvent(context.Context, *models.EventDTO) (uint, error)
	CreateBook(context.Context, uint, *models.BookDTO) (uint, error)
	ChangeBookStatus(context.Context, uint) error
	GetEvent(context.Context, uint) (*models.EventDB, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
