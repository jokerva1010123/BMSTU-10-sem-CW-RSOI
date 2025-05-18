package delivery

import (
	"RSOI_CW/internal/pkg/middleware"
	"RSOI_CW/internal/pkg/report"
	"net/http"
)

type ReportHandler struct {
	repo report.IRepo
}

func NewReportHandler(repo report.IRepo) *ReportHandler {
	return &ReportHandler{repo: repo}
}

func (h ReportHandler) CheckFilling(w http.ResponseWriter, r *http.Request) {
	body, status := h.repo.GetFilling()
	middleware.Response(w, status, body)
}

func (h ReportHandler) CheckByUsers(w http.ResponseWriter, r *http.Request) {
	body, status := h.repo.CheckByUsers()
	middleware.Response(w, status, body)
}
