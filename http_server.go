package main

import (
	"fmt"
	"net/http"
)

type HttpServer struct{}

func (hs *HttpServer) ServeHttp(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/ping" {
		hs.processPing(w)
	}
}

func (hs *HttpServer) processPing(w http.ResponseWriter) {
	fmt.Fprint(w, "pong")
}
