package raft

import (
	"fmt"
)

type Role int

const (
	Leader Role = iota
	Follower
	Candidate
)

type State struct {
	/*所有节点都需要持久保存以下3个字段*/
	CurrentTerm uint64 //第一次启动时初始化为0
	VotedFor    string
	Log         []Log

	/*所有节点只需在内存中保存以下2个字段*/
	CommitIndex uint64 //初始化为0
	LastApplied uint64 //初始化为0

	/*leader需要在内存中保存以下2个字段*/
	NextIndex  []uint64 //初始化为leader的最后一条日志索引+1
	MatchIndex []uint64 //leader已复制到其他节点的最高日志索引，初始化为0

	role Role
}

func (s *State) Leader() bool {
	return s.role == Leader
}

func (s *State) ChangeToLeader() {
	s.role = Leader
}

func (s *State) Follower() bool {
	return s.role == Follower
}

func (s *State) ChangeToFollower() {
	s.role = Follower
}

func (s *State) Candidate() bool {
	return s.role == Candidate
}

func (s *State) ChangeToCandidate() {
	s.role = Candidate
}

var NodeState State

func init() {
	NodeState = State{}
	fmt.Printf("init node, current term is %d\n", NodeState.CurrentTerm)
}
