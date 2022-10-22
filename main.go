package main

import (
	"fmt"
	"net"
	"os"

	"github.com/lanhuidong/raft/raft"
)

func main() {
	fmt.Println("Hello raft!")
	config := raft.ReadConfiguration()
	fmt.Printf("config: %v\n", config)
	listen, err := net.Listen("tcp", config.Endpoint())
	if err != nil {
		fmt.Println("start failed")
		os.Exit(1)
	}
	fmt.Printf("listen on: %s\n", config.Endpoint())
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(conn)
		}
	}
}
