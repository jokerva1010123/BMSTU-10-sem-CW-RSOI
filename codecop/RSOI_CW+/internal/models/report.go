package models

import "github.com/google/uuid"

type Report struct {
	TicketUUID uuid.UUID `json:"ticket_uuid"`
	FlightUUID uuid.UUID `json:"flight_uuid"`
	UserUUID   uuid.UUID `json:"user_uuid"`
	State      string    `json:"state,omitempty"`
}

type ReportFilling struct {
	FlightUUID uuid.UUID `json:"flight_uuid"`
	Tickets    uint      `json:"tickets"`
}

type ReportUsers struct {
	UserUUID    uuid.UUID `json:"user_uuid"`
	FlightsMade uint      `json:"flights_made"`
}
