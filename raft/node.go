package raft

type Role string

const (
	Follower  Role = "follower"
	Candidate Role = "candidate"
	Leader    Role = "leader"
)

type Node struct {
	ID           int
	Role         Role
	Term         int
	VotedFor     int // 0 means no vote in this term
	LastLogIndex int
	Votes        int
	votesFrom    map[int]bool
}

func NewNode(id, lastLogIndex int) *Node {
	return &Node{ID: id, Role: Follower, LastLogIndex: lastLogIndex}
}

// TODO: on a higher term, step down to follower, clear the old vote count,
// and make the node eligible for one vote in the new term. Never lower Term.
func (n *Node) ObserveTerm(term int) {
	_ = term
}

// TODO: grant at most one vote per term, and only to a candidate whose last
// log index is at least as fresh as this node's. Reject stale terms.
func (n *Node) RequestVote(term, candidateID, candidateLastLogIndex int) bool {
	_, _, _ = term, candidateID, candidateLastLogIndex
	return false
}
