package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteVoterResponse_Success(t *testing.T) {
	w := httptest.NewRecorder()

	expectedStatus := Status{
		Code:    200,
		Message: "Success",
	}

	writeVoterResponse(w, expectedStatus)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected content type %s, but got %s", "application/json", contentType)
	}

	var actualStatus Status
	err := json.Unmarshal(w.Body.Bytes(), &actualStatus)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	if actualStatus != expectedStatus {
		t.Errorf("Expected status %+v, but got %+v", expectedStatus, actualStatus)
	}
}

func TestWriteVoterResponse_Error(t *testing.T) {
	w := httptest.NewRecorder()

	invalidStatus := Status{
		Code:    500,
		Message: "Internal Server Error",
	}

	writeVoterResponse(w, invalidStatus)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected content type %s, but got %s", "application/json", contentType)
	}

	var actualStatus Status
	err := json.Unmarshal(w.Body.Bytes(), &actualStatus)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	if actualStatus != invalidStatus {
		t.Errorf("Expected status %+v, but got %+v", invalidStatus, actualStatus)
	}
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
