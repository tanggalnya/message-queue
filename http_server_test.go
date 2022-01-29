package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
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
	req, err := http.NewRequest("POST", "/guest-book/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GuestBookCreateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want  %v", status, http.StatusOK)
	}

	expected := `{"success": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
