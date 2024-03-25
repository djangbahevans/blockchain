package main

import "bytes"

// TXInput represents a transaction input.
type TXInput struct {
	TxId      []byte // The transaction ID that this input is referencing.
	Vout      int    // The index of the output being referenced.
	Signature []byte // The signature that proves the ownership of the output being spent.
	PubKey    []byte // The public key of the owner of the output being spent.
}

func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Equal(lockingHash, pubKeyHash)
}
