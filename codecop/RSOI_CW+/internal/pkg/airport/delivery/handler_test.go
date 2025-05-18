package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/airport"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAirport(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	id := uuid.New()

	mock := airport.NewMockIRepo(ctl)
	handler := NewAirportHandler(mock)

	mock.EXPECT().SelectAirport(id).Return(models.Airport{}, models.StatusOkey).Times(1)
	r := httptest.NewRequest("GET", "/airports/{UUID}", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"UUID": id.String(),
	})
	w := httptest.NewRecorder()

	handler.GetAirport(w, r)
}

func TestGetAirports(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mock := airport.NewMockIRepo(ctl)
	handler := NewAirportHandler(mock)

	mock.EXPECT().SelectAirports().Return([]models.Airport{}, models.StatusOkey).Times(1)
	r := httptest.NewRequest("GET", "/airports", strings.NewReader(fmt.Sprint()))

	w := httptest.NewRecorder()

	handler.GetAirports(w, r)
}
