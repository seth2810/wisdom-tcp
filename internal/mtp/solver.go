package mtp

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"

	"golang.org/x/crypto/argon2"
)

type Proof struct {
	Root  []byte `json:"root"`
	Nonce []byte `json:"nonce"`
}

func FindProof(cfg Config, challenge *Challenge) (*Proof, error) {
	// Generate memory-hard buffer using Argon2id
	salt, err := generateRandomBytes(challenge.SaltLength)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key salt: %w", err)
	}

	key := argon2.IDKey(
		challenge.Nonce,
		salt,
		cfg.TimeCost,
		challenge.MemorySize,
		cfg.Parallelism,
		32, // Output length
	)

	root, err := buildTree(key)
	if err != nil {
		return nil, fmt.Errorf("failed to build %q tree: %w", challenge.Nonce, err)
	}

	// Find nonce that satisfies difficulty
	nonce := make([]byte, 16)
	for {
		// Generate random nonce
		binary.BigEndian.PutUint64(nonce, uint64(time.Now().UnixNano()))

		// Create proof
		proof := &Proof{
			Root:  root,
			Nonce: nonce,
		}

		// Check if proof is valid
		if Verify(challenge, proof) {
			return proof, nil
		}
	}

}

func buildTree(key []byte) ([]byte, error) {
	// Create leaf nodes by splitting the key into 32-byte chunks
	leafCount := len(key) / 32
	leaves := make([][]byte, leafCount)
	for i := 0; i < leafCount; i++ {
		leaves[i] = key[i*32 : (i+1)*32]
	}

	// Build the Merkle tree
	for len(leaves) > 1 {
		// Create new level of nodes
		newLevel := make([][]byte, 0, (len(leaves)+1)/2)

		// Hash pairs of nodes
		for i := 0; i < len(leaves); i += 2 {
			var hash []byte
			if i+1 < len(leaves) {
				// Hash two nodes together
				combined := append(leaves[i], leaves[i+1]...)
				h := sha256.Sum256(combined)
				hash = h[:]
			} else {
				// Single node at the end, hash it with itself
				combined := append(leaves[i], leaves[i]...)
				h := sha256.Sum256(combined)
				hash = h[:]
			}
			newLevel = append(newLevel, hash)
		}

		leaves = newLevel
	}

	// Return the root hash
	return leaves[0], nil
}
