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
	msg := make(chan []byte)
	go func() {
		for _, node := range config.OtherNode() {
			go func(n raft.Node) {
				for {
					channel.ConnectToPeer(n, msg)
					time.Sleep(time.Duration(5) * time.Second)
				}
			}(node)
		}
	}()

	go func() {
		for data := range msg {
			fmt.Printf("receive: %s\n", string(data))
		}
	}()

	heartbeat := make(chan struct{})
	go func() {
		timer := time.NewTimer(5 * time.Second)
		select {
		case <-timer.C:
			fmt.Println("heartbeat timeout!")
			if raft.NodeState.Follower() {
				raft.NodeState.ChangeToCandidate()
				//发送选举消息
				//channel.SendMessage()
			}
		case <-heartbeat:
			fmt.Println("reset timeout")
			timer.Reset(5 * time.Second)
		}
	}()

	channel.ListenOnMessage(&config, msg)

}
