package blockchain

import (
	"fmt"
	"voting-system/utils"
)

type Block struct {
	Index        int
	Timestamp    string
	Transactions []*Transaction
	Proof        int
	PreviousHash string
}

// calculateHash calculates the hash of a block
func (b *Block) CalculateHash() string {
	// Convert transactions to a string
	transactionsString := ""
	for _, transaction := range b.Transactions {
		transactionsString += fmt.Sprintf("{Sender:%s, Recipient:%s, Amount:%d},", transaction.Sender, transaction.Recipient, transaction.Amount)
	}

	// Create the block data string
	blockData := fmt.Sprintf("%d%s%s%d%s", b.Index, b.Timestamp, transactionsString, b.Proof, b.PreviousHash)

	// Calculate and return the hash
	return utils.HashString(blockData)
}
