package main

type BlockChain struct {
	blocks []*Block
}

func NewGenesisBlock() *Block {
	return NewBlock([]byte("Genesis Block"), []byte{})
}

func NewBlockChain() *BlockChain {
	return &BlockChain{blocks: []*Block{NewGenesisBlock()}}
}

func (bc *BlockChain) AddBlock(data []byte) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.PrevBlockHash)
	bc.blocks = append(bc.blocks, newBlock)
}
