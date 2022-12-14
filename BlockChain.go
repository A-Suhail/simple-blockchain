package main

import (
	"crypto/sha256"
	"fmt"
)

func Hasher(ph []uint8, trans []string) []uint8 {
	hash := sha256.New()

	hash.Write([]byte(ph))
	for _, v := range trans {
		hash.Write([]byte(v))
	}

	return hash.Sum(nil)

}

// This is the Block struct
type Block struct {
	prevBlockHash []uint8
	blockHash     []uint8
	transactions  []string
}

// Constructor function for Block
func NewBlock(prevHash []uint8, trans []string) *Block {
	return &Block{
		prevBlockHash: prevHash,
		transactions:  trans,
		blockHash:     Hasher(prevHash, trans),
	}
}

func main() {
	//declare a blockchain with 0 blocks but capacity of 10 blocks
	blockchain := make([]Block, 0, 10)

	//this stores previous Block's hash
	var prevBlockH []uint8

	//create genesis block
	trans := []string{"satoshi gets 10 bt", "ivan gets 10 bt", "finney gets 10 bt"}
	var block Block = *NewBlock(prevBlockH, trans)
	prevBlockH = block.blockHash
	// fmt.Println("This is Genesis Block :")
	// fmt.Printf("prevHash : %v\nblockHash : %v\nTrasnsactions : %v\n", genesisBlk.prevBlockHash, genesisBlk.blockHash, genesisBlk.transactions)

	//push value to blockchain
	blockchain = append(blockchain, block)

	trans = []string{"ivan sent 2bt to satoshi", "finney sent 3bt to ivan"}
	block = *NewBlock(prevBlockH, trans)
	prevBlockH = block.blockHash

	blockchain = append(blockchain, block)

	trans = []string{"satoshi sent 6bt to ivan", "satoshi sent 6bt to finney"}
	block = *NewBlock(prevBlockH, trans)
	prevBlockH = block.blockHash

	blockchain = append(blockchain, block)

	for i, blk := range blockchain {
		fmt.Printf("This is Block : %v\n", i+1)
		fmt.Printf("prevHash : %x\nblockHash : %x\nTrasnsactions : %v\n", blk.prevBlockHash, blk.blockHash, blk.transactions)
	}
}
