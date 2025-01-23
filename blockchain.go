package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Index        int    // Position of the block in the chain
	Timestamp    string // Time when the block was created
	Data         string // Data stored in the block
	PreviousHash string // Hash of the previous block
	Hash         string // Hash of the current block
}

// Blockchain is a slice of blocks
var Blockchain []Block

// calculateHash computes the hash of a block
func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PreviousHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// generateBlock creates a new block
func generateBlock(previousBlock Block, data string) Block {
	newBlock := Block{
		Index:        previousBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Data:         data,
		PreviousHash: previousBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// isBlockValid checks the validity of a block
func isBlockValid(newBlock, previousBlock Block) bool {
	if previousBlock.Index+1 != newBlock.Index {
		return false
	}
	if previousBlock.Hash != newBlock.PreviousHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// main initializes the blockchain and adds new blocks
func main() {
	// Create the genesis block
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Data:         "Genesis Block",
		PreviousHash: "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	// Add more blocks
	newBlock := generateBlock(Blockchain[len(Blockchain)-1], "First Block Data")
	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		Blockchain = append(Blockchain, newBlock)
	}

	newBlock = generateBlock(Blockchain[len(Blockchain)-1], "Second Block Data")
	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		Blockchain = append(Blockchain, newBlock)
	}

	// Print the blockchain
	for _, block := range Blockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}
