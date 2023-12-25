package main

import (
	"bolt_version/lib/base58"
	"bolt_version/lib/ripemd160"
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

type Wallet struct {
	Private *ecdsa.PrivateKey
	PubKey  []byte
}

// 创建钱包

func NewWallet() *Wallet {
	curve := elliptic.P256()

	prinvateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}

	publicKey := prinvateKey.PublicKey

	pubKey := append(publicKey.X.Bytes(), publicKey.Y.Bytes()...)
	return &Wallet{Private: prinvateKey, PubKey: pubKey}
}

// 生成地址
func (W *Wallet) NewAddress() string {
	pubKey := W.PubKey
	rip160HashValue := HashPubKey(pubKey)
	version := byte(00)
	// 拼接version
	payload := append([]byte{version}, rip160HashValue...)
	checkCode := CheckSum(payload)
	payload = append(payload, checkCode...)
	// bctd库
	address := base58.Encode(payload)
	return address
}

func HashPubKey(data []byte) []byte {
	hash := sha256.Sum224(data)
	// 编码器
	rip160Hasher := ripemd160.New()

	_, err := rip160Hasher.Write(hash[:])
	if err != nil {
		panic(err)
	}

	rip160HashValue := rip160Hasher.Sum(nil)
	return rip160HashValue

}

func CheckSum(data []byte) []byte {
	hash1 := sha256.Sum256(data)
	// 前四个字节校验码
	hash2 := sha256.Sum256(hash1[:])
	checkCode := hash2[:4]
	// 25字节数据
	return checkCode
}
func IsValidAddress(address string) bool {
	//1. 解码
	addressByte := base58.Decode(address)

	if len(addressByte) < 4 {
		return false
	}

	//2. 取数据
	payload := addressByte[:len(addressByte)-4]
	checksum1 := addressByte[len(addressByte)-4:]

	//3. 做checksum函数
	checksum2 := CheckSum(payload)

	fmt.Printf("checksum1 : %x\n", checksum1)
	fmt.Printf("checksum2 : %x\n", checksum2)

	//4. 比较
	return bytes.Equal(checksum1, checksum2)
}
