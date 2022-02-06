package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"tanggalnya.com/message-queue/internal/services/rabbitMQ"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	queueName    = flag.String("queue-name", "events", "Ephemeral AMQP rabbitMQ name")
	exchange     = flag.String("exchange-name", "events_topic", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "topic", "Exchange type - direct|fanout|topic|x-custom")
	//body         = flag.String("body", "body test", "Body of message")
	reliable = flag.Bool("reliable", true, "Wait for the rabbitMQ confirmation before exiting")
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

func GuestBookCreateHandler(w http.ResponseWriter, r *http.Request) {
	sp := rabbitMQ.AmqpService{
		Uri: *uri,
	}
	p := rabbitMQ.NewAmqpChannel(sp)
	body := `
	{
		"data": {
		  "old": null,
		  "new": {
		    "from": null,
		    "is_public": true,
		    "name": "a",
		    "updated_at": "2022-01-29T16:14:55.062435+00:00",
		    "created_at": "2022-01-29T16:14:55.062435+00:00",
		    "id": "5f335ba6-39cb-466e-99ec-bbcb84b6cf85",
		    "message": "a",
		    "event_id": "9ebc9705-d6ab-4a82-8c7b-ba3802e5a241"
		  }
		}
	}`
	err := p.Publish(*queueName, *exchange, *exchangeType, body, *reliable)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"success": true}`)
}
