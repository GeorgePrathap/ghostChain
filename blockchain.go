package main

// Blockchain keeps a sequence of Blocks
type BlockChain struct {
	blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain
func (blockchain *BlockChain) AddBlock(data string) {
	prevBlock := blockchain.blocks[len(blockchain.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	blockchain.blocks = append(blockchain.blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
