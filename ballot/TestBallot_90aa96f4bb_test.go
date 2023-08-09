package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"testing"
)

func TestBallot(t *testing.T) {
	err := testSuccessfulBallotVote()
	if err != nil {
		t.Error(err)
	}

	err = testFailedBallotVote()
	if err != nil {
		t.Error(err)
	}
}

func testSuccessfulBallotVote() error {
	port := ""

	_, result, err := httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		return err
	}

	var initialRespData ResultBoard
	if err = json.Unmarshal(result, &initialRespData); err != nil {
		log.Printf("Failed to unmarshal get ballot response. %+v", err)
		return err
	}

	var ballotVoteReq Vote
	ballotVoteReq.CandidateID = fmt.Sprint(rand.Intn(10))
	ballotVoteReq.VoterID = fmt.Sprint(rand.Intn(10))
	reqBuff, err := json.Marshal(ballotVoteReq)
	if err != nil {
		log.Printf("Failed to marshall post ballot request %+v", err)
		return err
	}

	_, result, err = httpClientRequest(http.MethodPost, net.JoinHostPort("", port), "/", bytes.NewReader(reqBuff))
	if err != nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		return err
	}

	var postBallotResp Status
	if err = json.Unmarshal(result, &postBallotResp); err != nil {
		log.Printf("Failed to unmarshal post ballot response. %+v", err)
		return err
	}

	if postBallotResp.Code != 201 {
		return errors.New("post ballot resp status code")
	}

	_, result, err = httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		log.Printf("Failed to get final ballot count resp:%s error:%+v", string(result), err)
		return err
	}

	var finalRespData ResultBoard
	if err = json.Unmarshal(result, &finalRespData); err != nil {
		log.Printf("Failed to unmarshal get final ballot response. %+v", err)
		return err
	}

	if finalRespData.TotalVotes-initialRespData.TotalVotes != 1 {
		return errors.New("ballot vote count error")
	}

	return nil
}

func testFailedBallotVote() error {
	port := ""

	_, result, err := httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		return err
	}

	var initialRespData ResultBoard
	if err = json.Unmarshal(result, &initialRespData); err != nil {
		log.Printf("Failed to unmarshal get ballot response. %+v", err)
		return err
	}

	var ballotVoteReq Vote
	ballotVoteReq.CandidateID = "invalidCandidateID"
	ballotVoteReq.VoterID = fmt.Sprint(rand.Intn(10))
	reqBuff, err := json.Marshal(ballotVoteReq)
	if err != nil {
		log.Printf("Failed to marshall post ballot request %+v", err)
		return err
	}

	_, result, err = httpClientRequest(http.MethodPost, net.JoinHostPort("", port), "/", bytes.NewReader(reqBuff))
	if err != nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		return err
	}

	var postBallotResp Status
	if err = json.Unmarshal(result, &postBallotResp); err != nil {
		log.Printf("Failed to unmarshal post ballot response. %+v", err)
		return err
	}

	if postBallotResp.Code == 201 {
		return errors.New("expected failure for invalid candidate ID")
	}

	_, result, err = httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		log.Printf("Failed to get final ballot count resp:%s error:%+v", string(result), err)
		return err
	}

	var finalRespData ResultBoard
	if err = json.Unmarshal(result, &finalRespData); err != nil {
		log.Printf("Failed to unmarshal get final ballot response. %+v", err)
		return err
	}

	if finalRespData.TotalVotes != initialRespData.TotalVotes {
		return errors.New("ballot vote count error")
	}

	return nil
}
