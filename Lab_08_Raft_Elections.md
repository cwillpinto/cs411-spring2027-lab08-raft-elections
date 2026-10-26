# Lab 8: Raft Elections

**CS 411: Computing Paradigms - Spring 2027**  
**Week 10: March 15-21, 2027**  
**Due:** March 21, 2027  
**Value:** 100 points

## Purpose

Implement and trace Raft leader election in a five-node simulation. Preserve term monotonicity, one vote per term, majority election, and immediate step-down on a higher term.

## Learning objectives

- Trace follower candidate leader transitions
- Determine a majority from votes
- Preserve Raft term and vote invariants

## Before you begin

1. Review the Week 10 slide deck and assigned VoiceThread or reading.
2. Create your own repository from the Spring 2027 template linked in Blackboard. It contains a deterministic election-state model, public tests, and TODOs in `raft/node.go` and `raft/election.go`.
3. Run `go test ./... -run Election` before editing. Save the initial failures in `Lab8_Terminal.txt`.

## Part A Election trace (30 points)

Work one election by hand.

- In `Lab8_Trace.md` or PDF, trace a five-node election in which S4 starts term 1 and receives replies from S1 and S5. Record role, term, `VotedFor`, vote count, and last-log index at each state change.
- Show why S4's self-vote plus S1 and S5 is a strict majority of five. Then show why a duplicate reply from S1 cannot supply the third vote.
- Trace a split vote with no majority, then a new election in term 2. Use a fixed seed with the supplied `TimeoutFor` helper and report the timeout values so the scenario is reproducible.

## Part B Implement election state (45 points)

Complete the node state machine.

- Use the supplied seeded `TimeoutFor` helper; explain how distinct randomized timeouts reduce repeated split votes.
- Complete `StartElection`, `RequestVote`, `ReceiveVote`, and `ObserveTerm` while preserving term monotonicity and at most one candidate vote per term.
- Reject stale terms and candidates with a lower last-log index. Count each granted voter ID once and become leader only after a strict majority.
- Step down immediately when a higher term appears, including in a vote reply, and clear prior-term vote state.

## Part C Test (25 points)

Use deterministic seeds for reproducible elections.

- Run `go test ./... -run Election` and `go test -race ./...`; save complete commands/output in `Lab8_Terminal.txt`.
- Explain what each public test demonstrates: normal majority, stale-log rejection, duplicate vote handling, split vote followed by a new term, and leader step-down.
- State one behavior that the deterministic state model does not test, such as network delivery or persistent storage.

## Verification

- `go test ./... -run Election` and `go test -race ./...` pass.
- A candidate becomes leader only at three distinct votes in the five-node model; a stale-term reply cannot change the count.
- The trace contains term, voter ID, role, and log-freshness evidence.

## What to submit

Submit your own repository URL and:

1. Your own repository URL (created from the Spring 2027 template).
2. `raft/node.go`, `raft/election.go`, and the unchanged public `raft/election_test.go` in that repository.
3. `Lab8_Trace.pdf` or `Lab8_Trace.md`.
4. `Lab8_Terminal.txt`.

## Rubric

| Criterion | Points |
| --- | ---: |
| Election trace | 20 |
| Term and vote invariants | 30 |
| Majority and log freshness | 20 |
| Deterministic tests | 20 |
| Explanation | 10 |
| **Total** | **100** |

## AI usage and academic integrity

Add an **AI Usage** section to the required trace file. Name each tool and list the prompts or interactions that affected the work. If you did not use AI, write `AI Usage: None.` You must be able to explain and defend every submitted design and line of code.

## Late work

Late work may be accepted for up to three calendar days with a 25 percent deduction per day, following the course syllabus.

