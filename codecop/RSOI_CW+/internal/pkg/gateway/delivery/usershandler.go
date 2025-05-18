package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/middleware"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/streadway/amqp"
	"net/http"
	"os"
)

type UserHandler struct {
	ch  amqp.Channel
	que amqp.Queue
}

func NewUserHandler(ch amqp.Channel, que amqp.Queue) *UserHandler {
	return &UserHandler{ch: ch, que: que}
}

func (h *UserHandler) SendMessage(report models.Report) {

	body, _ := json.Marshal(report)

	_ = h.ch.Publish(
		"",         // exchange
		h.que.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}

func (h *UserHandler) IsAuthed(r *http.Request) bool {

	cookie := r.Header.Get("Authorization")
	if cookie == "" {
		return false
	}

	authAdrr, ok := os.LookupEnv("AUTH_URL")
	if !ok {
		panic("no  address")
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/verify", authAdrr),
		nil)
	if err != nil {
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

func (h *UserHandler) BuyTicket(w http.ResponseWriter, r *http.Request) {
	if !h.IsAuthed(r) {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	addr, ok := os.LookupEnv("TICKETS_URL")
	if !ok {
		panic("no address")
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/tickets", addr), nil)
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

	if resp.StatusCode == http.StatusOK {
		rp := &models.Report{}
		_ = easyjson.UnmarshalFromReader(r.Body, rp)
		rp.State = "Add"
		h.SendMessage(*rp)
	}

	middleware.CopyResponse(w, resp)
}

func (h *UserHandler) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	if !h.IsAuthed(r) {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	addr, ok := os.LookupEnv("TICKETS_URL")
	if !ok {
		panic("no address")
	}

	vars := mux.Vars(r)
	uuids, _ := vars["UUID"]
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/tickets/%s", addr, uuids), nil)
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

	if resp.StatusCode == http.StatusOK {
		rp := &models.Report{}
		_ = easyjson.UnmarshalFromReader(r.Body, rp)
		rp.State = "Delete"
		h.SendMessage(*rp)
	}

	middleware.CopyResponse(w, resp)
}

func (h *UserHandler) GetMiles(w http.ResponseWriter, r *http.Request) {
	if !h.IsAuthed(r) {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	addr, ok := os.LookupEnv("BONUS_URL")
	if !ok {
		panic("no address")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/miles", addr), nil)
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

func (h *UserHandler) GetTickets(w http.ResponseWriter, r *http.Request) {
	if !h.IsAuthed(r) {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	addr, ok := os.LookupEnv("TICKETS_URL")
	if !ok {
		panic("no address")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/tickets", addr), nil)
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
