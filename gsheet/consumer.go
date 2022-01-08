package gsheet

import (
	"fmt"
	"net/http"
)

type GsheetConsumer struct{}

func (g *GsheetConsumer) ServeHttp(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/ping" {
		g.processPing(w)
	}
}

func (g *GsheetConsumer) processPing(w http.ResponseWriter) {
	fmt.Fprint(w, "pong")
}
