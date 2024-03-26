package main

import "fmt"

func (cli *CLI) createBlockchain(address string) {
	if !ValidateAddress(address) {
		fmt.Println("ERROR: Address is not valid")
		return
	}

	bc := CreateBlockchain(address)
	defer bc.db.Close()

	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
