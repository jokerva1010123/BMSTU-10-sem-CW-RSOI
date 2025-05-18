package models

import "github.com/google/uuid"

type Airport struct {
	AirportUUID uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Description string    `json:"description"`
}
