package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"tanggalnya.com/message-queue/internal/models"
	"tanggalnya.com/message-queue/internal/services/publisher"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	queueName    = flag.String("queue-name", "events", "Ephemeral AMQP publisher name")
	exchange     = flag.String("exchange-name", "events_topic", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "topic", "Exchange type - direct|fanout|topic|x-custom")
	//body         = flag.String("body", "body test", "Body of message")
	reliable = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", HealthCheckHandler).Methods(http.MethodGet).Name("health_check")
	r.HandleFunc("/guest-book/create", GuestBookCreateHandler).Methods(http.MethodPost).Name("guest_book_create")

	log.Fatal(http.ListenAndServe("localhost:3000", r))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"ping": "pong"}`)
}

type DataPayload struct {
	Old *models.GuestBook
	New *models.GuestBook
}

type GuestBookPayload struct {
	Data DataPayload
}

func GuestBookCreateHandler(w http.ResponseWriter, r *http.Request) {
	var gbp GuestBookPayload
	err := json.NewDecoder(r.Body).Decode(&gbp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sp := publisher.AmqpService{
		Uri: *uri,
	}
	p := publisher.NewAmqpChannel(sp)
	body, err := json.Marshal(gbp.Data.New)
	err = p.Publish(*queueName, *exchange, *exchangeType, string(body), *reliable)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"success": true}`)
}
