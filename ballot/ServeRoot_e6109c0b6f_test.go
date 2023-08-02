package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServeRoot_e6109c0b6f(t *testing.T) {
	t.Run("GET request - success", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `{"votes": 0}`
		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("POST request - success", func(t *testing.T) {
		vote := Vote{
			VoterID:     "123",
			CandidateID: "456",
		}
		voteJSON, err := json.Marshal(vote)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(string(voteJSON)))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}

		expected := `{"code": 201, "message": "Vote saved sucessfully"}`
		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("POST request - invalid vote data", func(t *testing.T) {
		vote := Vote{
			VoterID:     "123",
			CandidateID: "",
		}
		voteJSON, err := json.Marshal(vote)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(string(voteJSON)))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := `{"code": 400, "message": "Vote is not valid. Vote can not be saved"}`
		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("Invalid request method", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPut, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

		expected := `{"code": 405, "message": "Bad Request. Vote can not be saved"}`
		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
