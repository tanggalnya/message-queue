package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	t.Run("return pong", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		response := httptest.NewRecorder()

		gsheet := &HttpServer{}
		gsheet.ServeHttp(response, request)

		assertStatus(t, response.Code, 200)
		assertResponseBody(t, response.Body.String(), "pong")
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
