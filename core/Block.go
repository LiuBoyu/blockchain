package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  //当前区块编号
	Timestamp     int64  //区块时间戳
	PrevBlockHash string //上一个区块哈希值
	Hash          string //当前区块哈希值
	Data          string //区块数据
}

func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

func GenerateNewBlock(prevBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevBlockHash = prevBlock.Hash
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}
func GenerateGenesisBlock() Block {
	prevBlock := Block{}
	prevBlock.Index = -1
	prevBlock.Hash = ""
	return GenerateNewBlock(prevBlock, "data:Genesis Block")

}
