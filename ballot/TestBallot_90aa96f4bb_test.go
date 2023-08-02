package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"testing"
)

func TestBallot(t *testing.T) {
	port := "8080"

	_, result, err := httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		t.Errorf("Failed to get ballot count resp:%s error:%+v", string(result), err)
	}

	var initalRespData ResultBoard
	if err = json.Unmarshal(result, &initalRespData); err != nil {
		t.Errorf("Failed to unmarshal get ballot response. %+v", err)
	}

	var ballotvotereq Vote
	ballotvotereq.CandidateID = fmt.Sprint(rand.Intn(10))
	ballotvotereq.VoterID = fmt.Sprint(rand.Intn(10))
	reqBuff, err := json.Marshal(ballotvotereq)
	if err != nil {
		t.Errorf("Failed to marshall post ballot request %+v", err)
	}

	_, result, err = httpClientRequest(http.MethodPost, net.JoinHostPort("", port), "/", bytes.NewReader(reqBuff))
	if err != nil {
		t.Errorf("Failed to get ballot count resp:%s error:%+v", string(result), err)
	}

	var postballotResp Status
	if err = json.Unmarshal(result, &postballotResp); err != nil {
		t.Errorf("Failed to unmarshal post ballot response. %+v", err)
	}

	if postballotResp.Code != 201 {
		t.Error("post ballot resp status code")
	}

	_, result, err = httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		t.Errorf("Failed to get final ballot count resp:%s error:%+v", string(result), err)
	}

	var finalRespData ResultBoard
	if err = json.Unmarshal(result, &finalRespData); err != nil {
		t.Errorf("Failed to unmarshal get final ballot response. %+v", err)
	}

	if finalRespData.TotalVotes-initalRespData.TotalVotes != 1 {
		t.Error("ballot vote count error")
	}
}

func httpClientRequest(method, url string, body io.Reader) (*http.Response, []byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, result, nil
}

type ResultBoard struct {
	TotalVotes int `json:"totalVotes"`
}

type Vote struct {
	CandidateID string `json:"candidateID"`
	VoterID     string `json:"voterID"`
}

type Status struct {
	Code int `json:"code"`
}
