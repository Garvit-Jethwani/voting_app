package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteCandidate_InvalidPayload(t *testing.T) {
	// Create a request with an invalid payload
	reqBody := []byte("invalid json")
	req, err := http.NewRequest("POST", "/deleteCandidate", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the deleteCandidate function with the invalid payload request
	deleteCandidate(rr, req)

	// Check the response status code
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
	}

	// Check the response body
	expectedResp := `{"status":400,"message":"Invalid request payload"}`
	if rr.Body.String() != expectedResp {
		t.Errorf("Expected response body %s, but got %s", expectedResp, rr.Body.String())
	}
}

func TestDeleteCandidate_CandidateNotFound(t *testing.T) {
	// Create a request with a valid payload
	reqBody := []byte(`{"name": "John Doe"}`)
	req, err := http.NewRequest("POST", "/deleteCandidate", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the deleteCandidate function with the valid payload request
	deleteCandidate(rr, req)

	// Check the response status code
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
	}

	// Check the response body
	expectedResp := `{"status":400,"message":"Candidate Not Found"}`
	if rr.Body.String() != expectedResp {
		t.Errorf("Expected response body %s, but got %s", expectedResp, rr.Body.String())
	}
}
