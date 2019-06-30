package core

import (
    "fmt"
    "log"
)

type Blockchain struct {
    Blocks []*Block
}

func NewBlockchain() *Blockchain {
    genesisBlock := GenerateGenesisBlock()
    blockchain := Blockchain{}
    blockchain.ApendBlock(&genesisBlock)
    return &blockchain
}

func (bc *Blockchain) SendData(data string) {
    preBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := GenerateNewBlock(*preBlock, data)
    bc.ApendBlock(&newBlock)

}

func (bc *Blockchain) ApendBlock(newBlock *Block) {
    if len(bc.Blocks) == 0 {
        bc.Blocks = append(bc.Blocks, newBlock)
        return
    }
    if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
        bc.Blocks = append(bc.Blocks, newBlock)
    } else {
        log.Fatal("v:invalid Block")
    }
}

func (bc *Blockchain) Print() {
    for _, block := range bc.Blocks {
        fmt.Printf("format:Index:%d\n", block.Index)
        fmt.Printf("format:Pre.Hash:%s\n", block.PrevBlockHash)
        fmt.Printf("format:Curr.Hash:%s\n", block.Hash)
        fmt.Printf("format:Data:%s\n", block.Data)
        fmt.Printf("format:Timestamp:%d\n", block.Timestamp)
    }
}

func isValid(newBlock Block, oldBlock Block) bool {
    if newBlock.Index-1 != oldBlock.Index {
        return false
    }
    if newBlock.PrevBlockHash != oldBlock.Hash {
        return false
    }
    if calculateHash(newBlock) != newBlock.Hash {
        return false
    }
    return true
}
