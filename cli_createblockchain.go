package main

import "fmt"

func (cli *CLI) createBlockchain(address string) {
	if !ValidateAddress(address) {
		fmt.Println("ERROR: Address is not valid")
		return
	}

	bc := CreateBlockchain(address)
	bc.db.Close()
	fmt.Println("Done!")
}
