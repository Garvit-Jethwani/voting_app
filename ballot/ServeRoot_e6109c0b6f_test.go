package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Vote struct {
	VoterID     string `json:"voter_id"`
	CandidateID string `json:"candidate_id"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func TestServeRoot_e6109c0b6f(t *testing.T) {
	t.Run("GET request - successful response", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
		}

		expectedResponse := `{"code":200,"message":"Success"}`
		if rr.Body.String() != expectedResponse {
			t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedResponse)
		}
	})

	t.Run("POST request - successful response", func(t *testing.T) {
		vote := Vote{
			VoterID:     "123456",
			CandidateID: "789012",
		}
		voteJSON, err := json.Marshal(vote)
		if err != nil {
			t.Fatalf("Failed to marshal vote: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(string(voteJSON)))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
		}

		expectedResponse := `{"code":201,"message":"Vote saved successfully"}`
		if rr.Body.String() != expectedResponse {
			t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedResponse)
		}
	})
}
