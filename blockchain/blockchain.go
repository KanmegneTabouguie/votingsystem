// blockchain/blockchain.go

package blockchain

import (
	"fmt"
	"voting-system/utils"
)

// Blockchain represents the main blockchain structure
type Blockchain struct {
	Blocks []*Block
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(transactions []*Transaction, proof int) {
	previousBlock := bc.getPreviousBlock()
	previousHash := ""
	if previousBlock != nil {
		previousHash = previousBlock.CalculateHash()
	}

	newBlock := &Block{
		Index:        len(bc.Blocks) + 1,
		Timestamp:    utils.GetCurrentTimestamp(),
		Transactions: transactions,
		Proof:        proof,
		PreviousHash: previousHash,
	}

	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain creates a new blockchain with the genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := &Block{
		Index:        1,
		Timestamp:    utils.GetCurrentTimestamp(),
		Transactions: []*Transaction{},
		Proof:        1,
		PreviousHash: "0", // Initial block has a hardcoded previous hash
	}

	return &Blockchain{Blocks: []*Block{genesisBlock}}
}

// getPreviousBlock returns the last block in the blockchain
func (bc *Blockchain) getPreviousBlock() *Block {
	if len(bc.Blocks) > 0 {
		return bc.Blocks[len(bc.Blocks)-1]
	}
	return nil
}

// CountVotes counts the votes for a specific candidate by ID
// CountVotes counts the votes for a specific candidate by ID
func (bc *Blockchain) CountVotes(candidateID string) int {
	voteCount := 0

	for _, block := range bc.Blocks {
		for _, transaction := range block.Transactions {
			// Assuming the recipient field in the transaction represents the candidate ID as a string
			if transaction.Recipient == fmt.Sprintf("Candidate %s", candidateID) {
				voteCount += transaction.Amount
			}
		}
	}

	return voteCount
}
