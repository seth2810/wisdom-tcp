package mtp

import "crypto/rand"

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// countLeadingZeroBits counts the number of leading zero bits in a byte slice
func countLeadingZeroBits(data []byte) uint8 {
	var count uint8

	for _, b := range data {
		if b == 0 {
			count += 8
		} else {
			// Count remaining zero bits in this byte
			for i := 7; i >= 0; i-- {
				if b&(1<<i) == 0 {
					count++
				} else {
					break
				}
			}

			break
		}
	}

	return count
}
