package main

import (
	"RSOI_CW/internal/pkg/bonus/delivery"
	"RSOI_CW/internal/pkg/bonus/repo"
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
	port, ok := os.LookupEnv("BONUS_PORT")

	if !ok {
		port = "8050"
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

	rp := repo.NewBonusRepo(*pool)
	handler := delivery.NewBonusHandler(rp)

	r := mux.NewRouter()
	r.Use(middleware.Cors)

	r.HandleFunc("/api/v1/miles", handler.GetBonus).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/miles", handler.AddBonus).Methods(http.MethodPost, http.MethodOptions)

	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)

	log.Print("bonus running on: ", srv.Addr)
	return srv.ListenAndServe()
}
