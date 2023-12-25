package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	//我们指定的难度值，现在是一个string类型，需要进行转换
	targetStr := "0000f00000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	//将难度值赋值给big.int，指定16进制的格式
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt
	return &pow
}

func (pow *ProofOfWork) Run() ([]byte, uint64) {
	//1. 拼装数据（区块的数据，还有不断变化的随机数）
	//2. 做哈希运算
	//3. 与pow中的target进行比较
	//a. 找到了，退出返回
	//b. 没找到，继续找，随机数加1

	var nonce uint64
	block := pow.block
	var hash [32]byte

	for {
		//1. 拼装数据（区块的数据，还有不断变化的随机数）
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}

		//将二维的切片数组链接起来，返回一个一维的切片
		blockInfo := bytes.Join(tmp, []byte{})

		hash = sha256.Sum256(blockInfo)
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		//比较当前的哈希与目标哈希值，如果当前的哈希值小于目标的哈希值，就说明找到了，否则继续找
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功！hash : %x, nonce : %d\n", hash, nonce)
			//break
			return hash[:], nonce
		} else {
			nonce++
		}

	}
}
