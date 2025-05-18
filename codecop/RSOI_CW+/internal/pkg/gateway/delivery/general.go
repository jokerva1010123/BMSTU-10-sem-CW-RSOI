package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type GeneralHandler struct {
}

func NewGeneralHandler() *GeneralHandler {
	return &GeneralHandler{}
}

func (h *GeneralHandler) GetFlights(w http.ResponseWriter, r *http.Request) {
	flightAdrr, ok := os.LookupEnv("FLIGHTS_URL")
	if !ok {
		panic("no address")
	}
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/flights", flightAdrr))
	if err != nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}
	defer resp.Body.Close()
	middleware.CopyResponse(w, resp)
}

func (h *GeneralHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	userAdrr, ok := os.LookupEnv("AUTH_URL")
	if !ok {
		panic("no address")
	}
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/users", userAdrr))
	if err != nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}
	defer resp.Body.Close()
	middleware.CopyResponse(w, resp)
}

func (h *GeneralHandler) GetAirports(w http.ResponseWriter, r *http.Request) {
	airAdrr, ok := os.LookupEnv("AIRPORTS_URL")
	if !ok {
		panic("no address")
	}
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/airports", airAdrr))
	if err != nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}
	defer resp.Body.Close()
	middleware.CopyResponse(w, resp)
}

func (h *GeneralHandler) GetAirport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuids, _ := vars["UUID"]

	airAdrr, ok := os.LookupEnv("AIRPORTS_URL")
	if !ok {
		panic("no address")
	}
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/airports/%s", airAdrr, uuids))
	if err != nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}
	defer resp.Body.Close()
	middleware.CopyResponse(w, resp)
}
