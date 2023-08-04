package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type BasicResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func writeBasicResponse(w http.ResponseWriter, resp *BasicResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)
	json.NewEncoder(w).Encode(resp)
}

func TestWriteBasicResponse_Success(t *testing.T) {
	recorder := httptest.NewRecorder()

	resp := &BasicResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	writeBasicResponse(recorder, resp)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected content type %s, but got %s", "application/json", contentType)
	}

	var parsedResp BasicResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &parsedResp)
	if err != nil {
		t.Errorf("Error unmarshaling response body: %v", err)
	}

	if parsedResp.Message != resp.Message {
		t.Errorf("Expected message %s, but got %s", resp.Message, parsedResp.Message)
	}

	if parsedResp.Data["key1"] != resp.Data["key1"] {
		t.Errorf("Expected data key1 value %s, but got %s", resp.Data["key1"], parsedResp.Data["key1"])
	}
	if parsedResp.Data["key2"] != resp.Data["key2"] {
		t.Errorf("Expected data key2 value %s, but got %s", resp.Data["key2"], parsedResp.Data["key2"])
	}
}

func TestWriteBasicResponse_Error(t *testing.T) {
	recorder := httptest.NewRecorder()

	resp := &BasicResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
	}

	writeBasicResponse(recorder, resp)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, recorder.Code)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected content type %s, but got %s", "application/json", contentType)
	}

	var parsedResp BasicResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &parsedResp)
	if err != nil {
		t.Errorf("Error unmarshaling response body: %v", err)
	}

	if parsedResp.Message != resp.Message {
		t.Errorf("Expected message %s, but got %s", resp.Message, parsedResp.Message)
	}
}
