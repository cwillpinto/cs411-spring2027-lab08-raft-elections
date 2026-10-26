package raft

import "math/rand"

// TimeoutFor returns a reproducible value in [150, 300) milliseconds. The
// seed belongs in the trace so a split vote can be replayed.
func TimeoutFor(seed int64, nodeID int) int {
	rng := rand.New(rand.NewSource(seed + int64(nodeID)*7919))
	return 150 + rng.Intn(150)
}

// TODO: begin a new term, vote for self, and enter candidate role.
func (n *Node) StartElection() {
}

// TODO: count each granted voter ID at most once and become leader only after
// a strict majority of totalNodes. Ignore stale replies and step down on a
// higher-term reply.
func (n *Node) ReceiveVote(replyTerm, voterID int, granted bool, totalNodes int) {
	_, _, _, _ = replyTerm, voterID, granted, totalNodes
}
