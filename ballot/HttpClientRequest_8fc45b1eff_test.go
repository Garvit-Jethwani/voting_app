package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func httpClientRequest(operation, hostAddr, command string, params io.Reader) (int, []byte, error) {

	url := "http://" + hostAddr + command
	if strings.Contains(hostAddr, "http://") {
		url = hostAddr + command
	}

	req, err := http.NewRequest(operation, url, params)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("Failed to create HTTP request." + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	defer resp.Body.Close()

	body, ioErr := ioutil.ReadAll(resp.Body)
	if hBit := resp.StatusCode / 100; hBit != 2 && hBit != 3 {
		if ioErr != nil {
			ioErr = fmt.Errorf("status code error %d", resp.StatusCode)
		}
	}
	return resp.StatusCode, body, ioErr
}

func TestHttpClientRequest_Success(t *testing.T) {
	operation := "GET"
	hostAddr := "example.com"
	command := "/api/resource"
	params := nil

	statusCode, body, err := httpClientRequest(operation, hostAddr, command, params)
	if err != nil {
		t.Error("Expected no error, but got:", err)
	}

	// TODO: Add assertions to validate the expected result
}

func TestHttpClientRequest_Failure(t *testing.T) {
	operation := "POST"
	hostAddr := "example.com"
	command := "/api/resource"
	params := nil

	statusCode, body, err := httpClientRequest(operation, hostAddr, command, params)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	// TODO: Add assertions to validate the expected result
}
