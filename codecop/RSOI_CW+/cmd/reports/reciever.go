package main

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/report/repo"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	if err := runReciever(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}

func runReciever() error {

	bdconn, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		return errors.New("no database url")
	}

	pool, err := pgxpool.Connect(context.Background(),
		bdconn)

	if err != nil {
		return err
	}

	rp := repo.NewReportRepo(*pool)

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

	q, err := ch.QueueDeclare(
		"ReportQueue", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	log.Printf(" [*] Welcome to my report service!")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			tk := &models.Report{}
			err = json.Unmarshal(d.Body, tk)
			if err == nil {
				rp.AddStat(*tk)
			}
			log.Printf("Received a report: %+v", tk)
		}
	}()
	<-forever

	return nil
}
