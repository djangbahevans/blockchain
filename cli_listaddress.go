package main

import (
	"fmt"
	"log"
)

func (cli *CLI) listAddresses(nodeId string) {
	wallets, err := NewWallets(nodeId)
	if err != nil {
		log.Panic(err)
	}

	addresses := wallets.GetAddresses()
	for _, address := range addresses {
		fmt.Println(address)
	}
}
