package models

type EventDTO struct {
	Name        string `json:"name" validate:"required"`
	SeatsNumber uint   `json:"seat_number" validate:"required"`
}

type BookDTO struct {
	PersonName string `json:"person_name" validate:"required"`
}

type EventDB struct {
	ID              uint   `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	SeatsNumberLeft uint   `json:"seat_number_left" validate:"required"`
}
