package main

import (
	"RSOI_CW/internal/pkg/airport/delivery"
	"RSOI_CW/internal/pkg/airport/repo"
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
	godotenv.Load(".env")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}

func run() error {
	port, ok := os.LookupEnv("AIRPORT_PORT")

	if !ok {
		port = "8020"
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

	repo := repo.NewAirportRepo(*pool)
	handler := delivery.NewAirportHandler(repo)

	r := mux.NewRouter()
	r.Use(middleware.Cors)
	//GET /airports/{airportUid}
	r.HandleFunc("/api/v1/airports/{UUID}",
		handler.GetAirport).Methods(http.MethodGet, http.MethodOptions)
	//GET /airports
	r.HandleFunc("/api/v1/airports",
		handler.GetAirports).Methods(http.MethodGet, http.MethodOptions)

	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)

	log.Print("airport running on: ", srv.Addr)
	return srv.ListenAndServe()
}
