package election

import "github.com/lanhuidong/raft/raft"

func VoteSelf(config *raft.Configuration) {
	raft.NodeState.VotedFor = config.SelfNode().Id
}
