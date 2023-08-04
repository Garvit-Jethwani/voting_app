package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeRoot_AllowOriginHeaderSet(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveRoot)

	handler.ServeHTTP(rr, req)

	if rr.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("Access-Control-Allow-Origin header is not set correctly")
	}
}

func TestServeRoot_AllowMethodsHeaderSet(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveRoot)

	handler.ServeHTTP(rr, req)

	if rr.Header().Get("Access-Control-Allow-Methods") != "POST, GET, OPTIONS, PUT, DELETE" {
		t.Error("Access-Control-Allow-Methods header is not set correctly")
	}
}
