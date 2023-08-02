package main

import (
	"sync"
	"testing"
)

func TestGetCandidatesVote_Success(t *testing.T) {
	candidateVotesStore = make(map[string]int)
	votes := getCandidatesVote()
	if len(votes) != 0 {
		t.Error("Expected empty candidateVotesStore, got non-empty")
	}
}

func TestGetCandidatesVote_ConcurrentAccess(t *testing.T) {
	candidateVotesStore = make(map[string]int)
	var wg sync.WaitGroup
	numGoroutines := 100
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			votes := getCandidatesVote()
			if len(votes) != 0 {
				t.Error("Expected empty candidateVotesStore, got non-empty")
			}
		}()
	}
	wg.Wait()
}
