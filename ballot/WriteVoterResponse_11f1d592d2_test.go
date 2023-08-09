package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Status struct {
	Message string `json:"message"`
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(status)
	if err != nil {
		log.Println("error marshaling response to vote request. error: ", err)
	}
	w.Write(resp)
}

func TestWriteVoterResponse_Success(t *testing.T) {
	recorder := httptest.NewRecorder()
	expectedStatus := Status{
		Message: "Success",
	}
	writeVoterResponse(recorder, expectedStatus)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
	expectedBody, _ := json.Marshal(expectedStatus)
	if recorder.Body.String() != string(expectedBody) {
		t.Errorf("expected response body %s, got %s", string(expectedBody), recorder.Body.String())
	}
}

func TestWriteVoterResponse_Error(t *testing.T) {
	recorder := httptest.NewRecorder()
	invalidStatus := Status{
		Message: "Invalid",
	}
	writeVoterResponse(recorder, invalidStatus)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
	expectedBody, _ := json.Marshal(invalidStatus)
	if recorder.Body.String() != string(expectedBody) {
		t.Errorf("expected response body %s, got %s", string(expectedBody), recorder.Body.String())
	}
}
