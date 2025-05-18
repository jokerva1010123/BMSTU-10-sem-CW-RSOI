package main

import (
	"RSOI_CW/internal/pkg/auth/delivery"
	"RSOI_CW/internal/pkg/auth/repo"
	"RSOI_CW/internal/pkg/middleware"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}

func run() error {
	port, ok := os.LookupEnv("AUTH_PORT")

	if !ok {
		port = "8000"
	}

	conn, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		return errors.New("no database url")
	}

	pool, err := pgxpool.Connect(context.Background(),
		conn)

	if err != nil {
		return err
	}

	repo := repo.NewAuthRepo(*pool)
	handler := delivery.NewAuthHandler(repo)

	r := mux.NewRouter()
	r.Use(middleware.Cors)

	//Авторизация.
	//header: Authorization: basic(<login>:<password>)
	//POST /auth -> JWT token
	r.HandleFunc("/api/v1/auth",
		handler.GetToken).Methods(http.MethodPost, http.MethodOptions)

	//Проверка токена пользователя.
	//header: Authorization: bearer <jwt>
	//POST /verify
	r.HandleFunc("/api/v1/verify",
		handler.CheckToken).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/v1/admin",
		handler.CheckAdminToken).Methods(http.MethodPost, http.MethodOptions)

	//Список всех пользователей. [A][G]
	//GET /users
	r.HandleFunc("/api/v1/users",
		handler.GetAllUsers).Methods(http.MethodGet, http.MethodOptions)

	//Добавление нового пользователя. [A][G]
	//POST /users
	r.HandleFunc("/api/v1/users",
		handler.AddUser).Methods(http.MethodPost, http.MethodOptions)

	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)

	log.Print("auth running on ", srv.Addr)
	return srv.ListenAndServe()
}
