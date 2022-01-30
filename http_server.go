package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"tanggalnya.com/message-queue/services"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	queueName    = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
	exchange     = flag.String("exchange", "test-exchange", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	body         = flag.String("body", "body test", "Body of message")
	reliable     = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", HealthCheckHandler).Methods(http.MethodGet).Name("health_check")
	r.HandleFunc("/guest-book/create", GuestBookCreateHandler).Methods(http.MethodPost).Name("guest_book_create")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"ping": "pong"}`)
}

func GuestBookCreateHandler(w http.ResponseWriter, r *http.Request) {
	publisher := services.Publisher{
		Uri:          *uri,
		QueueName:    *queueName,
		Exchange:     *exchange,
		ExchangeType: *exchangeType,
		Body:         *body,
		Reliable:     *reliable,
	}
	publisher.Publish()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"success": true}`)
}
