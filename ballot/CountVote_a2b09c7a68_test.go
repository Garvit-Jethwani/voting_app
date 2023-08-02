package main

import (
	"sort"
	"testing"
)

type CandidateVotes struct {
	CandidateID string
	Votes       int
}

type ResultBoard struct {
	Results     []CandidateVotes
	TotalVotes  int
}

func getCandidatesVote() map[string]int {
	// TODO: Implement this function
	return nil
}

func countVote() (ResultBoard, error) {
	votes := getCandidatesVote()
	res := ResultBoard{}
	for candidateID, votes := range votes {
		res.Results = append(res.Results, CandidateVotes{candidateID, votes})
		res.TotalVotes += votes
	}

	sort.Slice(res.Results, func(i, j int) bool {
		return res.Results[i].Votes > res.Results[j].Votes
	})
	return res, nil
}

func TestCountVote_Success(t *testing.T) {
	// TODO: Implement this test case
}

func TestCountVote_EmptyVotes(t *testing.T) {
	// TODO: Implement this test case
}
