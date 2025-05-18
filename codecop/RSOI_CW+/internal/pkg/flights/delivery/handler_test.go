package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/flights"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAllFlights(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mock := flights.NewMockIRepo(ctl)
	handler := NewFlightHandler(mock)

	mock.EXPECT().ReadFlights().Return([]models.Flight{}, models.StatusOkey).Times(1)
	r := httptest.NewRequest("GET", "/api/v1/flights", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	handler.AllFlights(w, r)
}

func TestUpdateFlight(t *testing.T) {

	flight := models.Flight{
		ID: uuid.New(),
		From:     uuid.UUID{},
		FromCity: "",
		To:       uuid.UUID{},
		ToCity:   "",
		Date:     time.Time{},
	}

	body, _ := json.Marshal(flight)
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mock := flights.NewMockIRepo(ctl)
	handler := NewFlightHandler(mock)

	mock.EXPECT().UpdateFlight(flight.ID, gomock.Any()).Return(models.StatusOkey).Times(1)
	r := httptest.NewRequest("UPDATE", "/api/v1/flights", strings.NewReader(string(body)))
	r = mux.SetURLVars(r, map[string]string{
		"UUID": flight.ID.String(),
	})
	w := httptest.NewRecorder()

	handler.UpdateFlight(w, r)
}

