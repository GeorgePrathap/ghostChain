package main

import (
	"time"
)

// Block keeps block headers
type Block struct {
	Timestamp                 int64
	Data, PrevBlockHash, Hash []byte
	Nonce                     int
}

// A function that takes a block and sets the hash of the block.
// func (block *Block) SetHash() {
// 	timeStamp := []byte(strconv.FormatInt(block.Timestamp, 10))
// 	header := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timeStamp}, []byte{})
// 	hash := sha256.Sum256(header)

// 	block.Hash = hash[:]
// }

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().UTC().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	proofOfWork := NewProofOfWork(block)
	nonce, hash := proofOfWork.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
