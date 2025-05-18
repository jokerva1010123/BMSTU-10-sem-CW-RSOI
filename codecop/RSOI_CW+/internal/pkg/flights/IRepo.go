package flights

import (
	"RSOI_CW/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source=IRepo.go -destination=IRepo_mock.go -package=flights
type IRepo interface {
	CreateFlight(flight models.Flight) int
	ReadFlight(FlightUUID uuid.UUID) (models.Flight, int)
	ReadFlights() ([]models.Flight, int)
	UpdateFlight(FlightUUID uuid.UUID, flight models.Flight) int
}
