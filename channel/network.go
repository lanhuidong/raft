package channel

import (
	"fmt"
	"net"
	"os"

	"github.com/lanhuidong/raft/raft"
)

func ListenOnMessage(config *raft.Configuration) {
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

func ConnectToPeer(n raft.Node) {
	conn, err := net.Dial("tcp", n.Endpoint())
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	defer conn.Close()
}
