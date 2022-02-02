package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HttpServiceTestSuite struct {
	suite.Suite
}

func (suite *HttpServiceTestSuite) SetupTest() {
}

func TestGuestBookCreateHandler(t *testing.T) {
	suite.Run(t, new(HttpServiceTestSuite))
}

func (suite *HttpServiceTestSuite) TestHealthCheckHandler() {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		suite.T().Errorf("handler returned wrong status code: got %v want  %v", status, http.StatusOK)
	}

	expected := `{"ping": "pong"}`
	if rr.Body.String() != expected {
		suite.T().Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
