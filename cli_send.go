package main

import "fmt"

func (cli *CLI) send(from, to string, amount int) {
	if !ValidateAddress(from) {
		fmt.Println("ERROR: Sender address is not valid")
		return
	}
	if !ValidateAddress(to) {
		fmt.Println("ERROR: Recipient address is not valid")
		return
	}
	bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	tx := NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := NewCoinbaseTx(from, "")
	txs := []*Transaction{cbTx, tx}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
	fmt.Println("Success!")
}
