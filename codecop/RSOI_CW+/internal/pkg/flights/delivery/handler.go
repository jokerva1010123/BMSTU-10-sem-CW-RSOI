package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/flights"
	"RSOI_CW/internal/pkg/middleware"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"net/http"
)

type FlightHandler struct {
	repo flights.IRepo
}

func NewFlightHandler(repo flights.IRepo) *FlightHandler {
	return &FlightHandler{repo: repo}
}

//список всех рейсов
//GET /flights
func (h *FlightHandler) AllFlights(w http.ResponseWriter, r *http.Request) {

	readFlights, status := h.repo.ReadFlights()
	middleware.Response(w, status, readFlights)
}

//добавить рейс
//POST /flights
func (h *FlightHandler) AddFlight(w http.ResponseWriter, r *http.Request) {
	flight := models.Flight{}
	err := easyjson.UnmarshalFromReader(
		r.Body,
		&flight,
	)
	if err != nil {
		middleware.Response(w, models.StatusBadUUID, nil)
		return
	}

	flight.ID = uuid.New()
	status := h.repo.CreateFlight(flight)
	middleware.Response(w, status, nil)

}

//изменить рейс
//PATCH /flights
func (h *FlightHandler) UpdateFlight(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uuids, _ := vars["UUID"]
	id, err := uuid.Parse(uuids)
	if err != nil {
		middleware.Response(w, models.StatusBadUUID, nil)
		return
	}
	flight := models.Flight{}
	err = easyjson.UnmarshalFromReader(
		r.Body,
		&flight,
	)
	if err != nil {
		middleware.Response(w, models.StatusBadUUID, nil)
		return
	}
	status := h.repo.UpdateFlight(id, flight)
	middleware.Response(w, status, nil)

}
