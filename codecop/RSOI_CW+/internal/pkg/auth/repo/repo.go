package repo

import (
	"RSOI_CW/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type AuthRepo struct {
	pool pgxpool.Pool
}

func NewAuthRepo(pool pgxpool.Pool) *AuthRepo {
	return &AuthRepo{pool: pool}
}

func (r *AuthRepo) GetUser(login string, password string) (models.User, int) {
	SelectUser := "SELECT \"UserUUID\", \"Login\", \"Role\" " +
		"FROM auth WHERE  \"Login\" = $1 AND \"Password\" = md5($2);"
	user := models.User{}

	row := r.pool.QueryRow(context.Background(), SelectUser, login, password)

	err := row.Scan(&user.UUID, &user.Login, &user.Role)
	if err != nil {
		log.Print(err)
		return models.User{}, models.StatusNotFound
	}

	return user, models.StatusOkey
}

func (r *AuthRepo) GetUsers() ([]models.User, int) {
	SelectUser := "SELECT \"UserUUID\", \"Login\", \"Role\" FROM auth;"

	result := []models.User{}

	rows, err := r.pool.Query(context.Background(), SelectUser)
	if err != nil {
		log.Print(err)
		return []models.User{}, models.StatusError
	}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.UUID, &user.Login, &user.Role)
		if err != nil {
			log.Print(err)
			return []models.User{}, models.StatusError
		}
		result = append(result, user)
	}

	return result, models.StatusOkey
}

func (r *AuthRepo) AddUser(user models.User) (models.User, int) {
	InsertUser := "INSERT INTO public.auth( \"UserUUID\", \"Login\", \"Password\", \"Role\")" +
		" VALUES ($1, $2, md5($3), $4);"
	if user.Role == "" {
		user.Role = "User"
	}
	user.UUID = uuid.New()

	exec, err := r.pool.Exec(context.Background(), InsertUser, user.UUID,
		user.Login, user.Password, user.Role)

	if err != nil {
		return models.User{}, models.StatusError
	}
	if exec.RowsAffected() == 0 {
		return models.User{}, models.StatusConflict
	}
	return user, models.StatusOkey
}
