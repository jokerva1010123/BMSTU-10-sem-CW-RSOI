package repo

import (
	"RSOI_CW/internal/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ReportRepo struct {
	pool pgxpool.Pool
}

func NewReportRepo(pool pgxpool.Pool) *ReportRepo {
	return &ReportRepo{pool: pool}
}

func (r *ReportRepo) AddStat(report models.Report) int {
	if report.State == "Add" {
		Insert := "INSERT INTO public.reports( \"UserUUID\", \"FlightUUID\", \"TicketUUID\")" +
			" VALUES ($1, $2, $3);"
		exec, err := r.pool.Exec(context.Background(),
			Insert, report.UserUUID, report.FlightUUID, report.TicketUUID)

		if err != nil {
			return models.StatusError
		}
		if exec.RowsAffected() == 0 {
			return models.StatusConflict
		}
		return models.StatusOkey
	} else {
		Delete := "DELETE FROM public.reports WHERE   \"UserUUID\" = $1" +
			" AND \"FlightUUID\" = $2 AND \"TicketUUID\" = $3;"
		exec, err := r.pool.Exec(context.Background(),
			Delete, report.UserUUID, report.FlightUUID, report.TicketUUID)

		if err != nil {
			return models.StatusError
		}
		if exec.RowsAffected() == 0 {
			return models.StatusNotFound
		}
		return models.StatusOkey
	}
}

func (r *ReportRepo) GetFilling() ([]models.ReportFilling, int) {
	Script := "SELECT \"FlightUUID\", COUNT(\"TicketUUID\") " +
		"FROM public.reports GROUP BY \"FlightUUID\";"

	query, err := r.pool.Query(context.Background(), Script)
	if err != nil {
		return nil, models.StatusError
	}
	result := []models.ReportFilling{}
	for query.Next() {
		rep := models.ReportFilling{}
		err = query.Scan(&rep.FlightUUID, &rep.Tickets)
		if err != nil {
			return nil, models.StatusError
		}
		result = append(result, rep)
	}
	return result, models.StatusOkey
}

func (r *ReportRepo) CheckByUsers() ([]models.ReportUsers, int) {
	Script := "SELECT \"UserUUID\", COUNT(\"FlightUUID\") FROM public.reports GROUP BY \"UserUUID\";"

	query, err := r.pool.Query(context.Background(), Script)
	if err != nil {
		return nil, models.StatusError
	}
	result := []models.ReportUsers{}
	for query.Next() {
		rep := models.ReportUsers{}
		err = query.Scan(&rep.UserUUID, &rep.FlightsMade)
		if err != nil {
			return nil, models.StatusError
		}
		result = append(result, rep)
	}
	return result, models.StatusOkey
}
