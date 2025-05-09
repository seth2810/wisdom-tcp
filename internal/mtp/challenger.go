package mtp

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
)

type Challenge struct {
	Nonce      []byte    `json:"nonce"`
	Timestamp  time.Time `json:"timestamp"`
	Difficulty uint8     `json:"difficulty"`
	MemorySize uint32    `json:"memory_size"`
	SaltLength uint32    `json:"salt_length"`
}

func GenerateNonce(difficulty uint8) []byte {
	nonce := make([]byte, 16)

	binary.BigEndian.PutUint64(nonce, uint64(time.Now().UnixNano()))

	return nonce
}

func Verify(challenge *Challenge, proof *Proof) bool {
	// Combine Merkle root and nonce
	data := append(proof.Root, proof.Nonce...)

	// Calculate hash
	hash := sha256.Sum256(data)

	// Check leading zero bits
	zeroBits := countLeadingZeroBits(hash[:])

	return zeroBits >= challenge.Difficulty
}
