package rpc

import (
	"github.com/lanhuidong/raft/raft"
)

type RequestVote struct {
	Term         uint64
	CandidateId  uint64
	LastLogIndex uint64
	LastLogTerm  uint64
}

type VoteResult struct {
	Term        uint64
	VoteGranted bool
}

type AppendEntries struct {
	Term         uint64
	LeaderId     uint64
	PrevLogIndex uint64
	PrevLogTerm  uint64
	Entries      []raft.Log
	LeaderCommit uint64
}

type Result struct {
	Term    uint64
	Success bool
}
