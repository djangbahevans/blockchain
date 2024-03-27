package main

import (
	"fmt"
	"log"
)

func (cli *CLI) startNode(nodeId, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeId)
	if len(minerAddress) > 0 {
		if ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}

	StartServer(nodeId, minerAddress)
}
