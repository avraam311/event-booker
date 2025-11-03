package events

import (
	"errors"

	"github.com/wb-go/wbf/dbpg"
)

var (
	ErrEventNotFound = errors.New("event not found")
	ErrNoSeatsOrEventNotFound = errors.New("no seats left or event not found")
	ErrBookNotFound  = errors.New("book not found")
)

type Repository struct {
	db *dbpg.DB
}

func NewRepository(db *dbpg.DB) *Repository {
	return &Repository{
		db: db,
	}
}
