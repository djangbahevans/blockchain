package main

import "fmt"

func (cli *CLI) reindexUTXO(nodeId string) {
	bc := NewBlockchain(nodeId)
	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
