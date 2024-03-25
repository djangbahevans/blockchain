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
	bc := NewBlockchain(from)
	defer bc.db.Close()

	tx := NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*Transaction{tx})
	fmt.Println("Success!")
}
