package unit1a

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

//定义挖矿的难度
const targetBits = 16

//防止计数器溢出
var maxNonce = math.MaxInt64

//工作量类型证明
type ProofofWork struct {
	block  *Block   //区块
	target *big.Int //目标值
}

//新建工作量证明
func NewProofofWork(b *Block) *ProofofWork {

	//创建一个大整数对象，初始值为1
	target := big.NewInt(1)

	//左移操作。得到一个目标操作
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofofWork{b, target}

	return pow
}

//准备工作量证明的计算数据
func (pow *ProofofWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			InttoHex(pow.block.Timestamp),
			InttoHex(int64(targetBits)), //挖矿难度
			InttoHex(int64(nonce)),      //计数器
		},
		[]byte{},
	)
	return data
}

//计算工作量证明
func (pow *ProofofWork) Run() (int, []byte) {
	// 1 用于保存Hash值的大整数
	var hashInt big.Int
	// 2 用于保存每次计算的Hash值的数组
	var hash [32]byte
	nonce := 0
	// 防止计数器整数溢出
	for nonce < maxNonce {
		// 3准备要计算的数据
		data := pow.prepareData(nonce)
		// 4求Hash值
		hash = sha256.Sum256(data)
		// 5打印Hash值
		//fmt.Printf("\r%x", hash)
		// 6 将Hash值设置为一个大整数
		hashInt.SetBytes(hash[:])
		// 7 与目标值比较
		if hashInt.Cmp(pow.target) == -1 {
			//fmt.Printf("Run()....\r%x", hash)
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

// 验证工作量证明的函数
func (pow *ProofofWork) Validate() bool {
	var hashInt big.Int
	// 准备数据
	data := pow.prepareData(pow.block.Nonce)
	// 计算Hash值
	hash := sha256.Sum256(data)
	// 将Hash值设置为大整数
	hashInt.SetBytes(hash[:])
	// 与目标值进行比较
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
