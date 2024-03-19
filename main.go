package main

import "fmt"

func main() {
	bc := NewBlockChain()

	bc.AddBlock([]byte("Send 1 btc to Ivan"))
	bc.AddBlock([]byte("Send 2 more BTC to Ivan"))

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
