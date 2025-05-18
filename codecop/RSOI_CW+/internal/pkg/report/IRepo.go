package report

import "RSOI_CW/internal/models"

//go:generate mockgen -source=IRepo.go -destination=IRepo_mock.go -package=report

type IRepo interface {
	AddStat(report models.Report) int
	GetFilling() ([]models.ReportFilling, int)
	CheckByUsers() ([]models.ReportUsers, int)
}
