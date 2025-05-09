package mtp

import "testing"

func TestCountLeadingZeroBits(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected uint8
	}{
		{
			name:     "All zeros",
			input:    []byte{0, 0, 0, 0},
			expected: 32,
		},
		{
			name:     "No leading zeros",
			input:    []byte{0xFF, 0xFF, 0xFF, 0xFF},
			expected: 0,
		},
		{
			name:     "Some leading zeros",
			input:    []byte{0x0F, 0xFF, 0xFF, 0xFF},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countLeadingZeroBits(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d leading zero bits, got %d", tt.expected, result)
			}
		})
	}
}
