package main

import "fmt"

func (cli *CLI) send(from, to string, amount int, nodeId string, mineNow bool) {
	if !ValidateAddress(from) {
		fmt.Println("ERROR: Sender address is not valid")
		return
	}
	if !ValidateAddress(to) {
		fmt.Println("ERROR: Recipient address is not valid")
		return
	}
	bc := NewBlockchain(nodeId)
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	wallets, err := NewWallets(nodeId)
	if err != nil {
		fmt.Println("ERROR: Wallets not found")
		return
	}
	wallet := wallets.GetWallet(from)

	tx := NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := NewCoinbaseTx(from, "")
		txs := []*Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		sendTx(knownNodes[0], tx)
	}
	fmt.Println("Success!")
}
