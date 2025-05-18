package models

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	TicketUUID uuid.UUID `json:"ticket_uuid,omitempty"`
	FlightUUID uuid.UUID `json:"flight_uuid"`
	UserUUID   uuid.UUID `json:"user_uuid,omitempty"`
	Date       time.Time `json:"date"`
}
