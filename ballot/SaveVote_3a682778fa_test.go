package main

import (
	"testing"
)

func TestSaveVote_Success(t *testing.T) {
	vote := Vote{
		CandidateID: "A",
	}

	err := saveVote(vote)

	if err != nil {
		t.Error("Expected nil error, got:", err)
	}

	// TODO: Check if vote count for candidate A has increased by 1
}

func TestSaveVote_Failure(t *testing.T) {
	vote := Vote{
		CandidateID: "B",
	}

	err := saveVote(vote)

	if err == nil {
		t.Error("Expected non-nil error, got:", err)
	}

	// TODO: Check if vote count for candidate B remains unchanged
}
