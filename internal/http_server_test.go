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

func (suite *HttpServiceTestSuite) TestGuestBookCreateHandler() {
	// TODO: fix test

	//payload := `
	//{
	//	"data": {
	//	  "old": null,
	//	  "new": {
	//	    "from": null,
	//	    "is_public": true,
	//	    "name": "a",
	//	    "updated_at": "2022-01-29T16:14:55.062435+00:00",
	//	    "created_at": "2022-01-29T16:14:55.062435+00:00",
	//	    "id": "5f335ba6-39cb-466e-99ec-bbcb84b6cf85",
	//	    "message": "a",
	//	    "event_id": "9ebc9705-d6ab-4a82-8c7b-ba3802e5a241"
	//	  }
	//	}
	//}`
	//
	//req, err := http.NewRequest("POST", "/guest-book/create", strings.NewReader(payload))
	//if err != nil {
	//	suite.T().Fatal(err)
	//}
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(GuestBookCreateHandler)
	//
	//mockAmqp := &mocks.AmqpChannel{}
	//fakeServer := server.NewServer(*uri)
	//fakeServer.Start()
	//
	//mockConn, err := amqptest.Dial(*uri)
	//mockAmqp.On("Channel", mock.Anything).Return(mockConn.Channel())
	//mockAmqp.On("Publish", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	//mockAmqp.EXPECT().Publish(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	//handler.ServeHTTP(rr, req)
	//
	//if status := rr.Code; status != http.StatusOK {
	//	suite.T().Errorf("handler returned wrong status code: got %v want  %v", status, http.StatusOK)
	//}
	//
	//expected := `{"success": true}`
	//if rr.Body.String() != expected {
	//	suite.T().Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	//}
}
