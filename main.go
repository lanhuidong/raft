package main

import (
	"fmt"
	"time"

	"github.com/lanhuidong/raft/channel"
	"github.com/lanhuidong/raft/raft"
)

func main() {
	fmt.Println("Hello raft!")
	config := raft.ReadConfiguration()
	fmt.Printf("config: %v\n", config)
	go func() {
		for _, node := range config.OtherNode() {
			go func(n raft.Node) {
				for {
					channel.ConnectToPeer(n)
					time.Sleep(time.Duration(5) * time.Second)
				}
			}(node)
		}
	}()

	channel.ListenOnMessage(&config)
}
