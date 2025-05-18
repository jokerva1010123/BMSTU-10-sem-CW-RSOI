package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/report"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCheckFilling(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := report.NewMockIRepo(ctl)
	handler := NewReportHandler(repo)

	r := httptest.NewRequest("GET", "/api/v1/flights", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	repo.EXPECT().GetFilling().Return([]models.ReportFilling{}, 0).Times(1)

	handler.CheckFilling(w, r)
}

func TestCheckByUsers(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := report.NewMockIRepo(ctl)
	handler := NewReportHandler(repo)

	r := httptest.NewRequest("GET", "/api/v1/flights", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	repo.EXPECT().CheckByUsers().Return([]models.ReportUsers{}, 0).Times(1)

	handler.CheckByUsers(w, r)
}
