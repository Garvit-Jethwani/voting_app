package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRunTest_13039fcb01(t *testing.T) {
	req1, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr1 := httptest.NewRecorder()
	handler1 := http.HandlerFunc(runTest)

	handler1.ServeHTTP(rr1, req1)

	if status := rr1.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedResponse1 := `{"Message":"Test Cases passed","Code":200}`
	if rr1.Body.String() != expectedResponse1 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr1.Body.String(), expectedResponse1)
	}

	req2, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(runTest)

	err = handler2.ServeHTTP(rr2, req2)

	if status := rr2.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expectedResponse2 := `{"Message":"Test Cases Failed with error : <error message>","Code":400}`
	if rr2.Body.String() != expectedResponse2 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr2.Body.String(), expectedResponse2)
	}
}
