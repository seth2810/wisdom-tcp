package mtp

import (
	"testing"

	"golang.org/x/crypto/argon2"
)

func TestBuildTree(t *testing.T) {
	key := argon2.IDKey(
		[]byte("test data"),
		nil,
		DefaultConfig.TimeCost,
		1024,
		DefaultConfig.Parallelism,
		32, // Output length
	)

	root, err := buildTree(key)
	if err != nil {
		t.Fatalf("Failed to generate Merkle tree: %v", err)
	}

	if len(root) != 32 {
		t.Errorf("Expected root length 32, got %d", len(root))
	}
}
