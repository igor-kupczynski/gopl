package main

import (
	"bufio"
	"bytes"
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
	// hash is the hash function to use, sha256 by default
	hash := func(data []byte) []byte {
		x := sha256.Sum256(data)
		return x[:]
	}

	// Allow the user to request different hash function
	for _, v := range os.Args[1:] {
		if v == "--sha384" {
			hash = func(data []byte) []byte {
				x := sha512.Sum384(data)
				return x[:]
			}
			break
		}
		if v == "--sha512" {
			hash = func(data []byte) []byte {
				x := sha512.Sum512(data)
				return x[:]
			}
			break
		}
	}

	// TODO: Is there a streaming sha256? we load everything to memory now
	var buf bytes.Buffer

	// Read stdio to buf
	w := bufio.NewWriter(&buf)
	_, err := io.Copy(w, os.Stdin)
	if err != nil {
		log.Fatalf("sha: %v", err)
	}
	err = w.Flush()
	if err != nil {
		log.Fatalf("sha: %v", err)
	}

	// Generate and print hash of buf
	sha := hash(buf.Bytes())
	fmt.Printf("%x\n", sha)
}
