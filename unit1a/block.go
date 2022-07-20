package unit1a

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
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
	pow := NewProofofWork(block)

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

//为区块类型增加序列化操作
func (b *Block) Serialize() []byte {

	//1,定义一个buff对象
	var result bytes.Buffer

	//2 新建一个go包的编码器
	encoder := gob.NewEncoder(&result)

	//3,执行编码
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	//4 调用buff对象bytes方法，返回一个字节切片
	return result.Bytes()
}

// 反序列化区块的函数
func DeserializeBlock(d []byte) *Block {

	//1,定义一个区块对象
	var block Block

	//2,创建一个解码器
	decorder := gob.NewDecoder(bytes.NewReader(d))

	//3 执行解码操作
	err := decorder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block

}
