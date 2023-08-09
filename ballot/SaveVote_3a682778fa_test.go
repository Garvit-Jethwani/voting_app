package main

import (
	"errors"
	"testing"
)

type Vote struct {
	CandidateID int
}

func saveVote(vote Vote) error {
	// TODO: Implement the saveVote function
	return nil
}

func TestSaveVote_Success(t *testing.T) {
	vote := Vote{
		CandidateID: 1,
	}

	err := saveVote(vote)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// TODO: Verify that the vote count for the candidate with ID 1 has increased by 1
}

func TestSaveVote_CandidateNotFound(t *testing.T) {
	vote := Vote{
		CandidateID: 999,
	}

	err := saveVote(vote)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	// TODO: Verify that the error message indicates that the candidate was not found
}

func TestSaveVote_NegativeCandidateID(t *testing.T) {
	vote := Vote{
		CandidateID: -1,
	}

	err := saveVote(vote)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	// TODO: Verify that the error message indicates that the candidate ID is invalid
}
