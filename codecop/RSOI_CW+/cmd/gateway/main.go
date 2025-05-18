package main

import (
	"RSOI_CW/internal/pkg/middleware"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"strings"

	"RSOI_CW/internal/pkg/gateway/delivery"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}

func run() error {
	port, ok := os.LookupEnv("GATEWAY_PORT")

	if !ok {
		port = "8000"
	}

	url, ok := os.LookupEnv("RABBIT_URL")

	if !ok {
		return errors.New("no rabbit url")
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	que, err := ch.QueueDeclare(
		"ReportQueue", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,
	)

	handler := delivery.NewGeneralHandler()
	userHandler := delivery.NewUserHandler(*ch, que)
	adminHandler := delivery.NewAdminHandler()
	r := mux.NewRouter()
	r.Use(middleware.InternalServerError)

	//G
	r.HandleFunc("/api/v1/users", handler.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/flights", handler.GetFlights).Methods("GET")
	r.HandleFunc("/api/v1/airports", handler.GetAirports).Methods("GET")
	r.HandleFunc("/api/v1/airports/{UUID}", handler.GetAirport).Methods("GET")

	//SG
	r.HandleFunc("/api/v1/tickets", userHandler.BuyTicket).Methods("POST")
	r.HandleFunc("/api/v1/tickets/{UUID}", userHandler.DeleteTicket).Methods("DELETE")
	r.HandleFunc("/api/v1/miles", userHandler.GetMiles).Methods("GET")
	r.HandleFunc("/api/v1/tickets", userHandler.GetTickets).Methods("GET")

	//SA
	r.HandleFunc("/api/v1/users", adminHandler.AddUser).Methods("POST")
	r.HandleFunc("/api/v1/flights", adminHandler.AddFlight).Methods("POST")
	r.HandleFunc("/api/v1/flights/{UUID}", adminHandler.EditFlight).Methods("PATCH")
	r.HandleFunc("/api/v1/reports/flights", adminHandler.GetFlightInfo).Methods("GET")
	r.HandleFunc("/api/v1/reports/flights-filling",
	adminHandler.GetInfo).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("GATEWAY_ORIGINS"), " "),
		AllowCredentials: true,
	})

	corsHandler := c.Handler(r)
	srv := http.Server{Handler: corsHandler, Addr: fmt.Sprintf(":%s", port)}

	log.Print("gateway running on ", srv.Addr)
	http.Handle("/", r)

	return srv.ListenAndServe()
}
