package repo

import (
	"RSOI_CW/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type BonusRepo struct {
	pool pgxpool.Pool
}

func NewBonusRepo(pool pgxpool.Pool) *BonusRepo {
	return &BonusRepo{pool: pool}
}

func (r *BonusRepo) GetBonus(UserUUID uuid.UUID) (int, int) {
	SelectBonus := "SELECT \"Balance\" FROM public.bonus WHERE WHERE \"UserUUID\"=$1;"
	row := r.pool.QueryRow(context.Background(), SelectBonus, UserUUID)

	var balance int
	err := row.Scan(&balance)
	if err != nil {
		return balance, models.StatusOkey
	}
	return 0, models.StatusNotFound
}

func (r *BonusRepo) SetBonus(UserUUID uuid.UUID, Extra int) (int, int) {
	UpdateBonus := "UPDATE public.bonus SET \"Balance\"=\"Balance\"+$1  WHERE \"UserUUID\"=$2 " +
		"RETURNING \"Balance\";"

	row := r.pool.QueryRow(context.Background(), UpdateBonus, Extra, UserUUID)

	var balance int
	err := row.Scan(&balance)
	if err != nil {
		return 0, models.StatusNotFound
	}
	return balance, models.StatusOkey
}

func (r *BonusRepo) CreateBonus(UserUUID uuid.UUID) int {
	InsertBonus := "INSERT INTO public.bonus(\"UserUUID\", \"Balance\") VALUES ($1, 0);"
	row := r.pool.QueryRow(context.Background(), InsertBonus, UserUUID)

	var balance int
	err := row.Scan(&balance)
	if err != nil {
		return models.StatusOkey
	}
	return models.StatusConflict
}
