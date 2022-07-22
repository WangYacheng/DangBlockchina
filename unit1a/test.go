package unit1a

import (
	"fmt"
)

func Show() {
	// 1 创建区块链对象
	bc := NewBlockchain()
	// 2 添加新的区块
	//bc.AddBlock("王小当发送1个比特币给廖璐瑶")
	//bc.AddBlock("廖璐瑶男朋友发送2个比特币给王小当")
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

func Show2() {
	// 1 新建区块链对象
	bc := NewBlockchain()
	// 2 关闭数据库
	defer bc.db.Close()
	// 3 创建命令行接口对象
	cli := CLI{bc}
	// 4 运行
	cli.Run()
}
