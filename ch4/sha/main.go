package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/igor-kupczynski/gopl/ch2/popcount"
	"io"
	"log"
	"os"
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

func main() {
	// hasher is the hasher function to use, sha256 by default
	hasher := sha256.New()

	// Allow the user to request different hasher function
	for _, v := range os.Args[1:] {
		if v == "--sha384" {
			hasher = sha512.New384()
		}
		if v == "--sha512" {
			hasher = sha512.New()
		}
	}

	// Stream stdio to the hasher
	if _, err := io.Copy(hasher, os.Stdin); err != nil {
		log.Fatalf("sha: %v", err)
	}

	fmt.Printf("%x\n", hasher.Sum(nil))
}
