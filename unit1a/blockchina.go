package unit1a

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {

	//获取数据中心的最后一个区块
	prevBlock := bc.blocks[len(bc.blocks)-1]

	//创建区块
	newBlock := NewBlock(data, prevBlock.Hash)

	//将新的区块添加到链上
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
