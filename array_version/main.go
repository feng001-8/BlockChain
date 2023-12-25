package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("123")
	bc.AddBlock("456")

	for i, block := range bc.blocks {

		fmt.Println("===========")
		fmt.Println("区块链长度:", i)
		fmt.Printf("前区块哈希值： %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值： %x\n", block.Hash)
		fmt.Printf("区块数据 :%s\n", block.Data)
	}
}
