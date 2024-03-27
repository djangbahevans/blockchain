package main

import "fmt"

func (cli *CLI) createWallet(nodeId string) {
	wallets, _ := NewWallets(nodeId)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeId)

	fmt.Printf("Your new address: %s\n", address)
}
