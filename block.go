package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// Block keeps block headers
type Block struct {
	Timestamp                 int64
	Data, PrevBlockHash, Hash []byte
	Nonce                     int
}

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

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
