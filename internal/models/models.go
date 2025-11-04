package models

import "time"

type EventDTO struct {
	Name        string `json:"name" validate:"required"`
	SeatsNumber uint   `json:"seats_number" validate:"required"`
}

type BookDTO struct {
	PersonName string `json:"person_name" validate:"required"`
}

type EventDB struct {
	ID              uint   `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	SeatsNumberLeft uint   `json:"seats_number_left" validate:"required"`
}

type BookDB struct {
	ID        uint      `validate:"required"`
	Book      string    `validate:"required"`
	CreatedAt time.Time `validate:"required"`
	EventID   uint      `validate:"required"`
}
