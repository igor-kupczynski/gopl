package main

import (
	"crypto/sha256"
	"github.com/igor-kupczynski/gopl/ch2/popcount"
)

// CountDifferentShaBits hashes its inputs and then counts the different bits
func CountDifferentBits(a []byte, b []byte) uint64 {
	shaA := sha256.Sum256(a)
	shaB := sha256.Sum256(b)

	diffBits := make([]byte, len(shaA))
	for i := range shaA {
		diffBits[i] = shaA[i] ^ shaB[i]
	}

	return popcount.PopCountByteSlice(diffBits)
}
