package channel

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/lanhuidong/raft/raft"
)

var id2ClientConn sync.Map

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
			go ReadFixedLengthMessage(conn)
		}
	}
}

func ReadFixedLengthMessage(conn net.Conn) {
	defer conn.Close()

	data := make([]byte, 0)
	buf := make([]byte, 1024)

	for {
		if n, err := conn.Read(buf); err == nil {
			var length uint32
			if n >= 4 && length == 0 {
				binary.Read(bytes.NewBuffer(buf[0:4]), binary.BigEndian, &length)
				data = append(data, buf[4:n]...)
			} else {
				data = append(data, buf[0:n]...)
			}
			if len(data) == int(length) {
				fmt.Printf("receive data %s\n", string(data))
			}
		}
	}
}

func ConnectToPeer(n raft.Node) {
	conn, err := net.Dial("tcp", n.Endpoint())
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	id2ClientConn.Store(n.Id, conn)
}

func SendMessage(n raft.Node, data []byte) {
	if conn, ok := id2ClientConn.Load(n.Id); ok {
		if n, err := conn.(net.Conn).Write(data); err == nil {
			fmt.Printf("data len %d, write %d\n", len(data), n)
		}
	}
}
