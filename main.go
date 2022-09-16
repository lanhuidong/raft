package main

import (
	"fmt"

	"github.com/lanhuidong/raft/raft"
)

func main() {
	fmt.Println("Hello raft!")
	config := raft.ReadConfiguration()
	fmt.Printf("config: %v\n", config)
}
