package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type AdminHandler struct {
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

func (h *AdminHandler) IsAdminAuthed(r *http.Request) bool {

	cookie := r.Header.Get("Authorization")
	if cookie == "" {
		return false
	}

	authAdrr, ok := os.LookupEnv("AUTH_URL")
	if !ok {
		panic("no  address")
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/admin", authAdrr),
		nil)

	if err != nil {
		log.Print(err)
		return false
	}
	req.Header.Set("Authorization", r.Header.Get("Authorization"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func (h *AdminHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	if !h.IsAdminAuthed(r) {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	authAdrr, ok := os.LookupEnv("AUTH_URL")
	if !ok {
		panic("no  address")
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/users", authAdrr),
		r.Body)

	if err != nil || req == nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}

	req.Header.Set("Authorization", r.Header.Get("Authorization"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}

	middleware.CopyResponse(w, resp)
}

func (h *AdminHandler) AddFlight(w http.ResponseWriter, r *http.Request) {
	if !h.IsAdminAuthed(r) {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	Adrr, ok := os.LookupEnv("FLIGHTS_URL")
	if !ok {
		panic("no  address")
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/flights", Adrr),
		r.Body)

	if err != nil || req == nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}

	req.Header.Set("Authorization", r.Header.Get("Authorization"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}

	middleware.CopyResponse(w, resp)
}

func (h *AdminHandler) EditFlight(w http.ResponseWriter, r *http.Request) {
	if !h.IsAdminAuthed(r) {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	Adrr, ok := os.LookupEnv("FLIGHTS_URL")
	if !ok {
		panic("no  address")
	}

	vars := mux.Vars(r)
	uuids, _ := vars["UUID"]

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/v1/flights/%s", Adrr, uuids),
		r.Body)

	if err != nil || req == nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}

	req.Header.Set("Authorization", r.Header.Get("Authorization"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}

	middleware.CopyResponse(w, resp)
}

func (h *AdminHandler) GetFlightInfo(w http.ResponseWriter, r *http.Request) {
	if !h.IsAdminAuthed(r) {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	addr, ok := os.LookupEnv("REPORTS_URL")
	if !ok {
		panic("no  address")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/reports/flights", addr),
		r.Body)

	if err != nil || req == nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}

	req.Header.Set("Authorization", r.Header.Get("Authorization"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}

	middleware.CopyResponse(w, resp)
}

func (h *AdminHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	if !h.IsAdminAuthed(r) {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	addr, ok := os.LookupEnv("REPORTS_URL")
	if !ok {
		panic("no  address")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/reports/flights-filling", addr),
		r.Body)

	if err != nil || req == nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}

	req.Header.Set("Authorization", r.Header.Get("Authorization"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		middleware.Response(w, models.StatusError, nil)
		return
	}

	middleware.CopyResponse(w, resp)
}
