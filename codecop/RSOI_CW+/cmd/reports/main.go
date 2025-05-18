package main

import (
	"RSOI_CW/internal/pkg/middleware"
	"RSOI_CW/internal/pkg/report/delivery"
	"RSOI_CW/internal/pkg/report/repo"
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
	port, ok := os.LookupEnv("REPORTS_PORT")

	if !ok {
		port = "8060"
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

	rp := repo.NewReportRepo(*pool)
	handler := delivery.NewReportHandler(rp)

	r := mux.NewRouter()
	r.Use(middleware.Cors)

	r.HandleFunc("/api/v1/reports/flights",
		handler.CheckByUsers).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/reports/flights-filling", handler.CheckFilling).
		Methods(http.MethodGet, http.MethodOptions)

	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)

	log.Print("report running on: ", srv.Addr)
	return srv.ListenAndServe()
}
