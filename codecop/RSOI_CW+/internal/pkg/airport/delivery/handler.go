package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/airport"
	"RSOI_CW/internal/pkg/middleware"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type AirportHandler struct {
	repo airport.IRepo
}

func NewAirportHandler(repo airport.IRepo) *AirportHandler {
	return &AirportHandler{repo: repo}
}

//GET /airports/{airportUid}
func (h *AirportHandler) GetAirport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuids, _ := vars["UUID"]
	id, err := uuid.Parse(uuids)
	if err != nil {
		middleware.Response(w, models.StatusBadUUID, nil)
		return
	} else {
		selectAirport, state := h.repo.SelectAirport(id)
		middleware.Response(w, state, selectAirport)
	}
}

//GET /airports
func (h *AirportHandler) GetAirports(w http.ResponseWriter, r *http.Request) {

	selectAirports, state := h.repo.SelectAirports()
	middleware.Response(w, state, selectAirports)
}
