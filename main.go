package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/lanhuidong/raft/raft"
)

func main() {
	fmt.Println("Hello raft!")
	config := raft.ReadConfiguration()
	fmt.Printf("config: %v\n", config)

	go func() {
		for _, node := range config.OtherNode() {
			go func(n *raft.Node) {
				for {
					conn, err := net.Dial("tcp", n.Endpoint())
					if err != nil {
						fmt.Printf("%s\n", err.Error())
						time.Sleep(time.Duration(5) * time.Second)
						continue
					}
					defer conn.Close()
					time.Sleep(time.Duration(5) * time.Second)
				}
			}(node)
		}
	}()

	listen, err := net.Listen("tcp", config.SelfNode().Endpoint())
	if err != nil {
		fmt.Println("start failed")
		os.Exit(1)
	}
	fmt.Printf("listen on: %s\n", config.SelfNode().Endpoint())
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(conn)
		}
	}
}
