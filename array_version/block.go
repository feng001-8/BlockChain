package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

// Block 定义结构
type Block struct {
	Version    uint64 //1.版本号
	PrevHash   []byte //2. 前区块哈希
	MerkelRoot []byte //3. Merkel根（
	TimeStamp  uint64 //4. 时间戳
	Difficulty uint64 //5. 难度值
	Nonce      uint64 //6. 随机数，也就是挖矿要找的数据
	Hash       []byte
	Data       []byte //b. 数据
}

// Uint64ToByte 实现一个辅助函数，功能是将uint64转成[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

// NewBlock 2. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}

	block.SetHash()

	return &block
}

// SetHash 3. 生成哈希
func (block *Block) SetHash() {
	//var blockInfo []byte
	//1. 拼装数据
	/*
		blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
		blockInfo = append(blockInfo, block.PrevHash...)
		blockInfo = append(blockInfo, block.MerkelRoot...)
		blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
		blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
		blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
		blockInfo = append(blockInfo, block.Data...)
	*/
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}

	//将二维的切片数组链接起来，返回一个一维的切片
	blockInfo := bytes.Join(tmp, []byte{})

	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
