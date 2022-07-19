package unit1a

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {

	//时间戳
	Timestamp int64

	//交易数据
	Data []byte

	//上一个区块哈希
	PrevBlockHash []byte

	//当前哈希
	Hash []byte

	//计数器
	Nonce int
}

//设置哈希
func (b *Block) SetHash() {
	timestamp := InttoHex(b.Timestamp)
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

//创建区块函数
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}

	//新建工作量证明
	pow := NewProofoWork(block)

	//计算工作量证明
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	fmt.Printf("NewBlock.Hash：%x\n", hash[:])
	return block
}

//创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
