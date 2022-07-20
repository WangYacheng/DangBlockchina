package unit1a

import (
	"fmt"
)

func Show() {
	// 1 创建区块链对象
	bc := NewBlockchain()
	// 2 添加新的区块
	bc.AddBlock("张三发送1个比特币给李四")
	bc.AddBlock("李四发送2个比特币给王五")
	// 3 遍历区块链
	// 3.1 创建迭代器
	bci := bc.Iterator()
	for {
		block := bci.Next()
		fmt.Printf("前一个区块的Hash值:%x\n", block.PrevBlockHash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("当前块Hash值:%x\n", block.Hash)
		//pow := NewProofofWork(block)
		//fmt.Println("PoW:", pow.Validate())
		fmt.Println()
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

}
