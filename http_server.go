package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"success": true}`)
}
