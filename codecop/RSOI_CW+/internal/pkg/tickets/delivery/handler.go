package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/middleware"
	"RSOI_CW/internal/pkg/tickets"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
)

type TicketHandler struct {
	repo tickets.IRepo
}

func NewTicketHandler(repo tickets.IRepo) *TicketHandler {
	return &TicketHandler{repo: repo}
}

//header: Authorization: bearer <jwt>
//GET /tickets
func (h *TicketHandler) GetMyTickets(w http.ResponseWriter, r *http.Request) {

	id := middleware.UserUUID(r)
	getTickets, status := h.repo.GetTickets(id)
	middleware.Response(w, status, getTickets)
}

//POST /tickets
//body: { flightUid, date }
func (h *TicketHandler) BuyTicket(w http.ResponseWriter, r *http.Request) {
	id := middleware.UserUUID(r)

	ticket := models.Ticket{}
	err := easyjson.UnmarshalFromReader(
		r.Body,
		&ticket,
	)

	if err != nil {
		log.Print(err)
		middleware.Response(w, models.StatusBadUUID, nil)
	} else {
		ticket.UserUUID = id
		ticket.TicketUUID = uuid.New()
		status := h.repo.CreateTicket(ticket)
		middleware.Response(w, status, nil)
	}
}

//header: Authorization: bearer <jwt>
//DELETE /tickets/{ticketUid}
func (h *TicketHandler) ReturnTicket(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uuids, _ := vars["UUID"]
	id, err := uuid.Parse(uuids)
	if err != nil {
		middleware.Response(w, models.StatusBadUUID, nil)
	} else {
		status := h.repo.DeleteTicket(id)
		middleware.Response(w, status, nil)
	}
}

//GET /tickets/{ticketUid}
func (h *TicketHandler) TicketInfo(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
	uuids, _ := vars["UUID"]
	id, err := uuid.Parse(uuids)
	if err != nil {
		middleware.Response(w, models.StatusBadUUID, nil)
	} else {
		ticket, status := h.repo.GetTicket(id)
		middleware.Response(w, status, ticket)
	}
}

//GET /tickets
func (h *TicketHandler) AllTicketInfo(w http.ResponseWriter, r *http.Request) {
	ticket, status := h.repo.GetAllTickets()
	middleware.Response(w, status, ticket)
}
