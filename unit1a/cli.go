package unit1a

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//命令行接口类型
type CLI struct {
	bc *Blockchain
}

//1 打印使用信息
func (cli *CLI) PrintUsage() {
	fmt.Print("使用信息")
	fmt.Println("addblock -data 区块信息-添加区块到区块链中")
	fmt.Println("printchain -打印区块链中所有区块信息")
}

//2 参数验证
func (cli *CLI) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.PrintUsage()
		os.Exit(1)
	}
}

//3 添加新区块
func (cli *CLI) addblock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}

// 4 打印区块信息
func (cli *CLI) printchain() {
	bci := cli.bc.Iterator()
	for {
		block := bci.Next()
		fmt.Printf("上一个区块的Hash值:%x\n", block.PrevBlockHash)
		fmt.Printf("区块信息:%s\n", block.Data)
		fmt.Printf("当前块Hash值:%x\n", block.Hash)
		//pow := NewProofofWork(block)
		//fmt.Println("Pow:", pow.Validate())
		fmt.Println()
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// 5 接受用户输入，解析并使用参数，调用相关函数
func (cli *CLI) Run() {
	cli.ValidateArgs()
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "区块信息")
	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		cli.printchain()
	default:
		cli.PrintUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addblock(*addBlockData)
	}
}
