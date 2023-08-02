package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBallot(t *testing.T) {
	// Create a new HTTP request with a GET method and a nil request body
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the runTest function with the response recorder and the request
	runTest(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := `{"message":"Test Cases passed","code":200}`
	if rr.Body.String() != expectedBody {
		t.Errorf("Expected response body %s, but got %s", expectedBody, rr.Body.String())
	}
}

func TestBallotFailure(t *testing.T) {
	// Create a new HTTP request with a POST method and a nil request body
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the runTest function with the response recorder and the request
	runTest(rr, req)

	// Check the response status code
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
	}

	// Check the response body
	expectedBody := `{"message":"Test Cases Failed with error : <error message>","code":400}`
	if rr.Body.String() != expectedBody {
		t.Errorf("Expected response body %s, but got %s", expectedBody, rr.Body.String())
	}
}

func runTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	log.Println("ballot endpoint tests running")
	status := Status{}
	err := TestBallot()
	if err != nil {
		status.Message = fmt.Sprintf("Test Cases Failed with error : %v", err)
		status.Code = http.StatusBadRequest
	}
	status.Message = "Test Cases passed"
	status.Code = http.StatusOK
	writeVoterResponse(w, status)
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.WriteHeader(status.Code)
	fmt.Fprintf(w, `{"message":"%s","code":%d}`, status.Message, status.Code)
}

type Status struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func TestBallot() error {
	// TODO: Implement your test cases here
	return nil
}
