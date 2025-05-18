package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/bonus"
	"RSOI_CW/internal/pkg/middleware"
	"net/http"
)

type BonusHandler struct {
	repo bonus.IRepo
}

func NewBonusHandler(repo bonus.IRepo) *BonusHandler {
	return &BonusHandler{repo: repo}
}

//получение бонусного баланса
func (h *BonusHandler) GetBonus(w http.ResponseWriter, r *http.Request) {
	id := middleware.UserUUID(r)
	balance, status := h.repo.GetBonus(id)
	bonus := models.Bonus{
		UserUUID: id,
		Balance:  balance,
	}
	middleware.Response(w, status, bonus)
}

//изменение бонусной программы
//здесь должен быть idПолета, чтобы было окей
func (h *BonusHandler) AddBonus(w http.ResponseWriter, r *http.Request) {
	userUUID := middleware.UserUUID(r)
	_, status := h.repo.SetBonus(userUUID, 1)
	middleware.Response(w, status, nil)

}

func (h *BonusHandler) NewBonusUser(w http.ResponseWriter, r *http.Request) {

	userUUID := middleware.UserUUID(r)
	status := h.repo.CreateBonus(userUUID)
	middleware.Response(w, status, nil)

}
