package raft

import "testing"

func TestElectionMajority(t *testing.T) {
	c := NewNode(4, 8)
	c.StartElection()
	if c.Role != Candidate || c.Term != 1 || c.VotedFor != 4 || c.Votes != 1 {
		t.Fatalf("candidate: %+v", c)
	}
	c.ReceiveVote(1, 1, true, 5)
	if c.Role == Leader {
		t.Fatal("two votes are not a majority of five")
	}
	c.ReceiveVote(1, 1, true, 5) // duplicate does not create a majority
	if c.Role == Leader {
		t.Fatal("duplicate vote counted twice")
	}
	c.ReceiveVote(1, 5, true, 5)
	if c.Role != Leader {
		t.Fatalf("three votes should elect leader: %+v", c)
	}
}

func TestElectionOneVoteAndLogFreshness(t *testing.T) {
	n := NewNode(2, 10)
	if n.RequestVote(1, 1, 9) {
		t.Fatal("stale log granted vote")
	}
	if !n.RequestVote(1, 3, 10) {
		t.Fatal("fresh candidate denied vote")
	}
	if n.RequestVote(1, 4, 11) {
		t.Fatal("second candidate granted vote in same term")
	}
	if !n.RequestVote(1, 3, 10) {
		t.Fatal("repeat request from same candidate should be stable")
	}
	if n.RequestVote(0, 3, 10) {
		t.Fatal("stale term granted vote")
	}
}

func TestElectionSplitThenHigherTerm(t *testing.T) {
	n := NewNode(1, 4)
	n.StartElection()
	n.ReceiveVote(1, 2, true, 5) // split: only two votes
	if n.Role == Leader {
		t.Fatal("split vote elected leader")
	}
	n.StartElection()
	if n.Term != 2 || n.Votes != 1 {
		t.Fatalf("new term: %+v", n)
	}
	n.ReceiveVote(1, 2, true, 5) // delayed reply from prior term
	if n.Votes != 1 {
		t.Fatal("stale reply counted")
	}
	n.ReceiveVote(2, 3, true, 5)
	n.ReceiveVote(2, 4, true, 5)
	if n.Role != Leader {
		t.Fatal("new-term majority did not elect leader")
	}
	n.ObserveTerm(3)
	if n.Role != Follower || n.Term != 3 || n.VotedFor != 0 || n.Votes != 0 {
		t.Fatalf("failed to step down: %+v", n)
	}
}

func TestElectionTimeoutsRepeatAndVary(t *testing.T) {
	if TimeoutFor(42, 1) != TimeoutFor(42, 1) {
		t.Fatal("timeout is not deterministic")
	}
	seen := map[int]bool{}
	for id := 1; id <= 5; id++ {
		seen[TimeoutFor(42, id)] = true
	}
	if len(seen) < 2 {
		t.Fatal("all nodes have same timeout")
	}
}
