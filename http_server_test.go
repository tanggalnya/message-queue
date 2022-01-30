package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NeowayLabs/wabbit/amqptest/server"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want  %v", status, http.StatusOK)
	}

	expected := `{"ping": "pong"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGuestBookInsertHandler(t *testing.T) {
	jsonStr := []byte(`
		{
    	  "event": {
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
    	  }
    	}`)
	t.Run("when correct request body it return ok", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/guest-book/create", bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GuestBookCreateHandler)

		handler.ServeHTTP(rr, req)

		assertResponse(t, http.StatusOK, rr.Code)

		expected := `{"success": true}`
		assertResponseBody(t, rr.Body.String(), expected)
	})

	t.Run("when data valid it call publish rabbitmq event", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/guest-book/create", bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GuestBookCreateHandler)
		handler.ServeHTTP(rr, req)

		fakeServer := server.NewServer("amqp://localhost:5672/%2f")
		fakeServer.Start()
	})
}

func assertResponse(t *testing.T, expected int, actual int) {
	t.Helper()
	if expected != actual {
		t.Errorf("handler returned wrong status code: got %v want %v", actual, expected)
	}

}

func assertResponseBody(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
