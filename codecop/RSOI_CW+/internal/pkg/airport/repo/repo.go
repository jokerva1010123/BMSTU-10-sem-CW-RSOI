package repo

import (
	"RSOI_CW/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type AirportRepo struct {
	pool pgxpool.Pool
}

func NewAirportRepo(pool pgxpool.Pool) *AirportRepo {
	return &AirportRepo{pool: pool}
}

func (r *AirportRepo) SelectAirports() ([]models.Airport, int) {
	SelectAll := "SELECT \"AirportUUID\", \"AirportName\", \"City\" FROM public.airport;"

	row, err := r.pool.Query(context.Background(), SelectAll)
	if err != nil {
		return []models.Airport{}, models.StatusError
	}

	result := []models.Airport{}
	for row.Next() {
		airport := models.Airport{}
		if err := row.Scan(&airport.AirportUUID, &airport.Name, &airport.City); err != nil {
			log.Print(err)
			return []models.Airport{}, models.StatusError
		}
		result = append(result, airport)
	}
	return result, models.StatusOkey
}

func (r *AirportRepo) SelectAirport(uuid uuid.UUID) (models.Airport, int) {
	SelectOne := "SELECT  \"AirportUUID\", \"AirportName\", \"City\", \"Description\" FROM public.airport" +
		" WHERE \"AirportUUID\" = $1;"

	row := r.pool.QueryRow(context.Background(), SelectOne, uuid)

	airport := models.Airport{}

	if err := row.Scan(&airport.AirportUUID, &airport.Name, &airport.City, &airport.Description); err != nil {
		return models.Airport{}, models.StatusError
	}
	return airport, models.StatusOkey
}
