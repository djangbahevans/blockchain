package main

import "bytes"

// TXOutput represents a transaction output.
type TXOutput struct {
	Value      int    // The value of the output.
	PubKeyHash []byte // The hash of the public key associated with the output.
}

func NewTXOutput(value int, address string) *TXOutput {
	txo := TXOutput{value, nil}
	txo.Lock([]byte(address))

	return &txo
}

func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}
