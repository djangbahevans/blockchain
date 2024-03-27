package main

import (
	"fmt"
	"strconv"
)

func (cli *CLI) printChain(nodeId string) {
	bc := NewBlockchain(nodeId)
	defer bc.db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Print("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
