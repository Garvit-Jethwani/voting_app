package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllCandidates_ef98167bf1(t *testing.T) {
	req, err := http.NewRequest("GET", "/candidates", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	getAllCandidates(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", status, http.StatusOK)
	}
}
