package unit1a

import (
	"fmt"
	"strconv"
)

func Show() {
	// 1 创建区块链对象
	bc := NewBlockchain()
	// 2 添加新的区块
	bc.AddBlock("张三发送1个比特币给李四")
	bc.AddBlock("李四发送2个比特币给王五")
	// 3 遍历区块链

	for _, block := range bc.blocks {
		fmt.Printf("前一个区块的Hash值:%x\n", block.PrevBlockHash)
		fmt.Printf("交易数据:%s\n", block.Data)
		fmt.Printf("当前块的Hash值：%x\n", block.Hash)

		pow := NewProofoWork(block)
		fmt.Printf("PoW：%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
