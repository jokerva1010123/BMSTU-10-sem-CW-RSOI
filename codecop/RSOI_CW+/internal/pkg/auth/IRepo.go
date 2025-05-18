package auth

import "RSOI_CW/internal/models"

//go:generate mockgen -source=IRepo.go -destination=IRepo_mock.go -package=auth

type IRepo interface {
	GetUser(login string, password string) (models.User, int)
	GetUsers() ([]models.User, int)
	AddUser(user models.User) (models.User, int)
}
