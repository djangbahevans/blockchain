package main

func main() {
	bc := NewBlockChain()
	defer bc.db.Close()

	// bc.AddBlock([]byte("Send 1 btc to Ivan"))
	// bc.AddBlock([]byte("Send 2 more BTC to Ivan"))

	cli := CLI{bc}
	cli.Run()
}
