package unit1a

import (
	"fmt"
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

			fmt.Println("这里是创世区块")
		} else {
			//
			tip = b.Get([]byte("1"))
			fmt.Println("这里不是创世区块")
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
		lastHash = b.Get([]byte("1"))
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
		err = b.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash
		fmt.Printf("AddBlock()"+data+"：%x\n", bc.tip[:])
		lastHash := b.Get([]byte("1"))
		fmt.Printf("lastHash= :%x\n", lastHash[:])
		return nil
	})
}

// 创建区块链迭代器的方法
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}
	return bci
}

// test
func Test1() {
	db, err := bolt.Open(dbFile, 0600, nil)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash := b.Get([]byte("1"))
		result := b.Get(lastHash[:])

		//result := b.Get([]byte("1"))
		block := DeserializeBlock(result)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

func Test2() {
	db, err := bolt.Open(dbFile, 0600, nil)
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		c := bucket.Cursor()

		/*
			index := 0
			for {
				k, v := c.Next()
				if k == nil || v == nil {
					return nil
				}
				index = index + 1
				fmt.Println(index)
			}
		*/

		for k, v := c.First(); k != nil; k, v = c.Next() {
			str := string(k)
			if str == "1" {
				fmt.Println("这里是1")
				result := bucket.Get(v)
				block := DeserializeBlock(result)
				fmt.Printf("Data:%s\n", block.Data)
				fmt.Printf("Hash:%x\n", block.Hash)

			} else {
				fmt.Println("这里是 链")
				block := DeserializeBlock(v)
				fmt.Printf("Data:%s\n", block.Data)
				fmt.Printf("Hash:%x\n", block.Hash)

			}
			fmt.Println()
		}

		return nil

	})

	if err != nil {
		log.Panic(err)
	}

}

func Test3() {}
