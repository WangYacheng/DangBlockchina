package unit1a

import (
	"log"

	"github.com/boltdb/bolt"
)

//数据库文件
const dbFile = "blockchain.db"

//表名
const blockBucket = "blocks"

type Blockchain struct {
	tip []byte //最后一个区块的哈希
	db  *bolt.DB
}

func NewBlockchain() *Blockchain {
	var tip []byte
	//新建或打开 数据库文件
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	//2,执行读写事务
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		if b == nil {
			//2.1 创建创世区块
			genesis := NewGenesisBlock()

			//2.2 创建表
			b, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic(err)
			}

			//2.3 写入创世块
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			//2.4 写入最后一个区块的Hash值
			err = b.Put([]byte("1"), genesis.Hash)
			tip = b.Get([]byte("1"))
		} else {
			//
			tip = b.Get([]byte("1"))
		}
		//fmt.Println("NewBlockchain")
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	//创建区块链对象
	bc := Blockchain{tip, db}
	return &bc
}

// 为区块链类型添加方法AddBlock，添加新区块
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte
	// 1 从数据库中读取最后一个区块的Hash值
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	// 2 创建新的区块
	newBlock := NewBlock(data, lastHash)
	// 3 将新的区块存储到数据库中
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		// 3.1 写入新的区块
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		// 3.2 写入最后一个区块的Hash
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash
		return nil
	})
}

// 创建区块链迭代器的方法
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}
	return bci
}
