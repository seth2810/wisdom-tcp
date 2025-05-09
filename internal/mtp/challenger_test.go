package mtp

import (
	"testing"
)

func TestGenerateNonce(t *testing.T) {
	difficulty := uint8(4)
	nonce := GenerateNonce(difficulty)

	if nonce == nil {
		t.Fatal("Generated challenge is nil")
	}

	if len(nonce) != 16 {
		t.Errorf("Expected nonce length 16, got %d", len(nonce))
	}
}

func TestVerifyProof(t *testing.T) {
	difficulty := uint8(4)
	challenge := &Challenge{
		Nonce:      GenerateNonce(difficulty),
		Difficulty: difficulty,
	}

	// Create a proof that should fail
	proof := &Proof{
		Root:  make([]byte, 32),
		Nonce: make([]byte, 16),
	}

	if Verify(challenge, proof) {
		t.Error("Expected invalid proof to fail verification")
	}

	// TODO: Add test for valid proof once Merkle tree generation is implemented
}
