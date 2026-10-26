# CS411 Spring 2027 · Lab 8: Raft Elections

Starter template: https://github.com/cwillpinto/cs411-spring2027-lab08-raft-elections

Read the included handout. Complete the TODOs in `raft/node.go` and `raft/election.go`. This is a deterministic election-state model: the public tests send vote requests and replies directly instead of introducing sockets or real timers. `TimeoutFor` provides reproducible randomized timeout values for the trace.

```sh
go test ./... -run Election
go test -race ./...
```

The unchanged starter compiles but fails the election invariant tests. Public tests are representative, not exhaustive.

## Work in your own GitHub repository

The instructor starter is an upstream source. Keep your completed work in a repository under your own GitHub account, then submit the repository URL and the evidence listed in the assignment handout through Blackboard.

If Blackboard links a template repository, choose **Use this template**, create the new repository under your account, and clone that new repository. Check `git remote -v` before pushing; `origin` must be your repository.

For a direct clone, first create an empty repository under your account. Replace the placeholders below with the starter link supplied in Blackboard and your own empty repository URL:

```sh
git clone https://github.com/cwillpinto/cs411-spring2027-lab08-raft-elections
cd <cloned-directory>
git remote rename origin upstream
git remote add origin <your-empty-repository-url>
git remote -v
git push -u origin HEAD
```

Commit and push your completed work to your own `origin`. Do not push to the instructor's upstream repository. Use the exact submission files and verification commands in the assignment handout.
