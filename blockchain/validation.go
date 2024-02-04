// blockchain/validation.go

package blockchain

import (
	"fmt"
	"strings"
	"voting-system/utils"
)

// IsValid checks the integrity and validity of the blockchain
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		// Check if the previous hash matches
		if currentBlock.PreviousHash != previousBlock.CalculateHash() {
			return false
		}

		// Check if the proof of work is valid
		if !isValidProof(previousBlock.Proof, currentBlock.Proof, previousBlock.CalculateHash()) {
			return false
		}
	}

	return true
}

// isValidProof checks if the given proof is valid
func isValidProof(lastProof, proof int, previousHash string) bool {
	// Hash the combination of the last proof, current proof, and previous hash
	guess := fmt.Sprintf("%d%d%s", lastProof, proof, previousHash)
	guessHash := utils.HashString(guess)

	// Check if the hash has a specific number of leading zeros (adjust as needed)
	return strings.HasPrefix(guessHash, "0000")
}
