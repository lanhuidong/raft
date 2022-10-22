package raft

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Cluster struct {
	Id    string
	Nodes map[string]Node
}

type Node struct {
	Id      string `json:"id"`
	Address string `json:"address"`
	Port    uint32 `json:"port"`
}

func (node Node) Endpoint() string {
	return node.Address + ":" + strconv.Itoa(int(node.Port))
}

type Configuration struct {
	Id    string `json:"id"`
	Nodes []Node `json:"nodes"`
}

func (c Configuration) SelfNode() *Node {
	for _, val := range c.Nodes {
		if c.Id == val.Id {
			return &val
		}
	}
	return nil
}

func (c Configuration) OtherNode() []*Node {
	nodes := make([]*Node, 0, len(c.Nodes)-1)
	for _, val := range c.Nodes {
		if c.Id != val.Id {
			nodes = append(nodes, &val)
		}
	}
	return nodes
}

func ReadConfiguration() Configuration {
	configFile, err := os.Open("conf/cluster.json")
	if err != nil {
		fmt.Printf("open config file error: %v", err)
		os.Exit(0)
	}
	defer configFile.Close()

	fileContent, err := io.ReadAll(configFile)
	if err != nil {
		fmt.Printf("read config file error: %v", err)
		os.Exit(0)
	}

	var config Configuration
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		fmt.Printf("config file invalid: %v", err)
		os.Exit(0)
	}
	return config
}
