package main

import (
	"sync"
	"testing"
)

var once sync.Once
var candidateVotesStore map[string]int

func getCandidatesVote() map[string]int {
	once.Do(func() {
		candidateVotesStore = make(map[string]int)
	})
	return candidateVotesStore
}

func TestGetCandidatesVote_Success(t *testing.T) {
	candidateVotesStore = nil
	getCandidatesVote()
	if candidateVotesStore == nil {
		t.Error("candidateVotesStore is not initialized")
	}
	candidateVotesStore["Candidate A"] = 10
	candidateVotesStore["Candidate B"] = 5
	if candidateVotesStore["Candidate A"] != 10 {
		t.Error("Failed to add votes for Candidate A")
	}
	if candidateVotesStore["Candidate B"] != 5 {
		t.Error("Failed to add votes for Candidate B")
	}
}

func TestGetCandidatesVote_AlreadyInitialized(t *testing.T) {
	candidateVotesStore = nil
	candidateVotesStore = make(map[string]int)
	getCandidatesVote()
	if len(candidateVotesStore) != 0 {
		t.Error("candidateVotesStore should not be re-initialized")
	}
}
