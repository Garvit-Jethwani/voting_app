package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddCandidate(t *testing.T) {
	// Test case 1: Valid request payload
	reqBody := `{"name": "John Doe", "imageUrl": "https://example.com/johndoe.jpg"}`
	req := httptest.NewRequest(http.MethodPost, "/candidates", strings.NewReader(reqBody))
	rec := httptest.NewRecorder()

	addCandidate(rec, req)

	resp := rec.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	// Test case 2: Invalid request payload
	reqBody = `{"name": "John Doe"}`
	req = httptest.NewRequest(http.MethodPost, "/candidates", strings.NewReader(reqBody))
	rec = httptest.NewRecorder()

	addCandidate(rec, req)

	resp = rec.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, resp.StatusCode)
	}
}
