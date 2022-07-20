package unit1a

import (
	"log"

	"github.com/boltdb/bolt"
)

// 定义迭代器类型
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// 为迭代器类型添加获取下一个区块的方法
func (i *BlockchainIterator) Next() *Block {
	var block *Block
	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock((encodedBlock))
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	i.currentHash = block.PrevBlockHash
	return block
}
