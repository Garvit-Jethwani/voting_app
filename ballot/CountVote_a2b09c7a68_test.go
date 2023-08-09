package main

import (
	"sort"
	"testing"
)

type ResultBoard struct {
	Results    []CandidateVotes
	TotalVotes int
}

type CandidateVotes struct {
	CandidateID string
	Votes       int
}

func getCandidatesVote() map[string]int {
	// TODO: Implement this function to fetch candidates' votes from a data source
	return map[string]int{
		"candidate1": 10,
		"candidate2": 15,
		"candidate3": 5,
	}
}

func countVote() (res ResultBoard, err error) {
	votes := getCandidatesVote()
	for candidateID, votes := range votes {
		res.Results = append(res.Results, CandidateVotes{candidateID, votes})
		res.TotalVotes += votes
	}

	sort.Slice(res.Results, func(i, j int) bool {
		return res.Results[i].Votes > res.Results[j].Votes
	})
	return res, err
}

func TestCountVote_Success(t *testing.T) {
	expectedResult := ResultBoard{
		Results: []CandidateVotes{
			{"candidate2", 15},
			{"candidate1", 10},
			{"candidate3", 5},
		},
		TotalVotes: 30,
	}

	result, err := countVote()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(result.Results) != len(expectedResult.Results) {
		t.Errorf("Expected %d results, but got %d", len(expectedResult.Results), len(result.Results))
	}

	for i := 0; i < len(result.Results); i++ {
		if result.Results[i].CandidateID != expectedResult.Results[i].CandidateID {
			t.Errorf("Expected candidate ID %s, but got %s", expectedResult.Results[i].CandidateID, result.Results[i].CandidateID)
		}
		if result.Results[i].Votes != expectedResult.Results[i].Votes {
			t.Errorf("Expected votes %d, but got %d", expectedResult.Results[i].Votes, result.Results[i].Votes)
		}
	}

	if result.TotalVotes != expectedResult.TotalVotes {
		t.Errorf("Expected total votes %d, but got %d", expectedResult.TotalVotes, result.TotalVotes)
	}
}

func TestCountVote_EmptyVotes(t *testing.T) {
	getCandidatesVote = func() map[string]int {
		return map[string]int{}
	}

	expectedResult := ResultBoard{
		Results:    []CandidateVotes{},
		TotalVotes: 0,
	}

	result, err := countVote()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(result.Results) != len(expectedResult.Results) {
		t.Errorf("Expected %d results, but got %d", len(expectedResult.Results), len(result.Results))
	}

	for i := 0; i < len(result.Results); i++ {
		if result.Results[i].CandidateID != expectedResult.Results[i].CandidateID {
			t.Errorf("Expected candidate ID %s, but got %s", expectedResult.Results[i].CandidateID, result.Results[i].CandidateID)
		}
		if result.Results[i].Votes != expectedResult.Results[i].Votes {
			t.Errorf("Expected votes %d, but got %d", expectedResult.Results[i].Votes, result.Results[i].Votes)
		}
	}

	if result.TotalVotes != expectedResult.TotalVotes {
		t.Errorf("Expected total votes %d, but got %d", expectedResult.TotalVotes, result.TotalVotes)
	}
}
