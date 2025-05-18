package airport

import (
	"RSOI_CW/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source=IRepo.go -destination=IRepo_mock.go -package=airport

type IRepo interface {
	SelectAirports() ([]models.Airport, int)
	SelectAirport(uuid uuid.UUID) (models.Airport, int)
}
