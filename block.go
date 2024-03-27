package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

// Block represents a block in the blockchain.
type Block struct {
	Timestamp     int64          // The timestamp when the block was created.
	Transactions  []*Transaction // The list of transactions included in the block.
	PrevBlockHash []byte         // The hash of the previous block in the blockchain.
	Hash          []byte         // The hash of the current block.
	Nonce         int            // The nonce value used in mining the block.
	Height        int
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte, height int) *Block {
	block := Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0, height}
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return &block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}, 0)
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		return []byte{}
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewBuffer(d))
	err := decoder.Decode(&block)
	if err != nil {
		return &block
	}

	return &block
}

func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := NewMerkleTree(transactions)

	return mTree.RootNode.Data
}
