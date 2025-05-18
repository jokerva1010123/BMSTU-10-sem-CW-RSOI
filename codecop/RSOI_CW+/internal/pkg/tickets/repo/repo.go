package repo

import (
	"RSOI_CW/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type TicketRepo struct {
	pool pgxpool.Pool
}

func NewTicketRepo(pool pgxpool.Pool) *TicketRepo {
	return &TicketRepo{pool: pool}
}

func (r *TicketRepo) GetTickets(uuid uuid.UUID) ([]models.Ticket, int) {
	SelectMy := "SELECT \"TicketUUID\", \"FlightUUID\", \"UserUUID\",  \"Date\" " +
		"FROM public.ticket WHERE \"UserUUID\" = $1;"

	log.Print(uuid)
	rows, err := r.pool.Query(context.Background(), SelectMy, uuid)
	if err != nil {
		log.Print(err)
		return nil, models.StatusError
	}

	result := []models.Ticket{}
	for rows.Next() {
		ticket := models.Ticket{}
		err = rows.Scan(&ticket.TicketUUID, &ticket.FlightUUID, &ticket.UserUUID, &ticket.Date)
		if err != nil {
			log.Print(err)
			return nil, models.StatusError
		}
		result = append(result, ticket)
	}

	return result, models.StatusOkey
}

func (r *TicketRepo) GetAllTickets() ([]models.Ticket, int) {
	SelectMy := "SELECT \"TicketUUID\", \"FlightUUID\", \"UserUUID\" " +
		"FROM public.ticket;"

	rows, err := r.pool.Query(context.Background(), SelectMy)
	if err != nil {
		return nil, models.StatusError
	}

	result := []models.Ticket{}
	for rows.Next() {
		ticket := models.Ticket{}
		err = rows.Scan(&ticket.TicketUUID, &ticket.FlightUUID, &ticket.UserUUID)
		if err != nil {
			return nil, models.StatusError
		}
		result = append(result, ticket)
	}

	return result, models.StatusOkey
}

func (r *TicketRepo) GetTicket(uuid uuid.UUID) (models.Ticket, int) {
	SelectMy := "SELECT \"TicketUUID\", \"FlightUUID\", \"UserUUID\" " +
		"FROM public.ticket WHERE \"TicketUUID\" = $1;"

	rows := r.pool.QueryRow(context.Background(), SelectMy, uuid)
	ticket := models.Ticket{}
	err := rows.Scan(&ticket.TicketUUID, &ticket.FlightUUID, &ticket.UserUUID)
	if err != nil {
		return ticket, models.StatusError
	}

	return ticket, models.StatusOkey
}

func (r *TicketRepo) DeleteTicket(uuid uuid.UUID) int {
	DeleteTicket := "DELETE FROM public.ticket WHERE \"TicketUUID\" = $1;"

	exec, err := r.pool.Exec(context.Background(), DeleteTicket, uuid)
	if err != nil {
		return models.StatusError
	}

	if exec.RowsAffected() == 0 {
		return models.StatusNotFound
	}

	return models.StatusOkey
}

func (r *TicketRepo) CreateTicket(ticket models.Ticket) int {
	InsertTicket := "INSERT INTO public.ticket(\"TicketUUID\", \"FlightUUID\", \"UserUUID\"," +
		" \"Date\" ) " +
		" VALUES ($1, $2, $3, $4);"

	ticket.TicketUUID = uuid.New()
	exec, err := r.pool.Exec(context.Background(), InsertTicket,
		ticket.TicketUUID, ticket.FlightUUID, ticket.UserUUID, ticket.Date)

	if err != nil {
		return models.StatusError
	}

	if exec.RowsAffected() == 0 {
		return models.StatusConflict
	}

	return models.StatusOkey
}
