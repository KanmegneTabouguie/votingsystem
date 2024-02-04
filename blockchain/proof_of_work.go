// blockchain/proof_of_work.go

package blockchain

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

// ProofOfWork represents the proof-of-work algorithm
type ProofOfWork struct {
	TargetBits int
	Nonce      int
	BlockHash  string
	MaxNonce   int
}

// NewProofOfWork initializes a new proof-of-work
func NewProofOfWork(targetBits int, blockHash string) *ProofOfWork {
	return &ProofOfWork{
		TargetBits: targetBits,
		Nonce:      0,
		BlockHash:  blockHash,
		MaxNonce:   1 << 32,
	}
}

// Run performs the proof-of-work algorithm
func (pow *ProofOfWork) Run() (int, string) {
	var hashInt big.Int
	var hash [32]byte

	target := big.NewInt(1)
	target.Lsh(target, uint(256-pow.TargetBits))

	for pow.Nonce < pow.MaxNonce {
		data := pow.prepareData(pow.Nonce)
		hash = sha256.Sum256([]byte(data))

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			break
		} else {
			pow.Nonce++
		}
	}

	return pow.Nonce, fmt.Sprintf("%x", hash[:])
}

// Validate checks if the generated hash meets the target difficulty
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.Nonce)
	hash := sha256.Sum256([]byte(data))

	hashInt.SetBytes(hash[:])

	target := big.NewInt(1)
	target.Lsh(target, uint(256-pow.TargetBits))

	return hashInt.Cmp(target) == -1
}

// prepareData prepares the data for hashing by combining block information and the nonce
func (pow *ProofOfWork) prepareData(nonce int) string {
	return fmt.Sprintf("%s%d%d", pow.BlockHash, nonce, pow.TargetBits)
}
