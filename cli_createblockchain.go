package main

import "fmt"

func (cli *CLI) createBlockChain(address string) {
	if !ValidateAddress(address) {
		fmt.Println("ERROR: Address is not valid")
		return
	}

	bc := CreateBlockChain(address)
	bc.db.Close()
	fmt.Println("Done!")
}
