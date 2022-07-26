package unit1a

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

const subsidy = 10 //挖矿奖励的比特币

type Transaction struct {
	ID   []byte
	vin  []TXInput
	vout []TXOutput
}

//创币交易
func NewCoinbaseTx(to, data string) *Transaction {}

//计算交易的Hash的值，并保存到交易的ID字段中
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}
