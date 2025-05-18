package bonus

import "github.com/google/uuid"

//go:generate mockgen -source=IRepo.go -destination=IRepo_mock.go -package=bonus

type IRepo interface {
	GetBonus(UserUUID uuid.UUID) (int, int)
	SetBonus(UserUUID uuid.UUID, Extra int) (int, int)
	CreateBonus(UserUUID uuid.UUID) int
}
