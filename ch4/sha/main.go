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

	// It pays off to stream. We have this 71M input:
	//
	//		$ ll -h | grep input
	//		-rw-rw-r-- 1 igor igor  71M Mar 21 15:28 input.bin
	//
	// If we load it to mem, the program will eat up 200M+ of memory:
	//
	// 		$ cat input.bin  | time -v ./sha-nostream
	// 		e017fa5f3b78104b18c9e3ec00a678e513095cd4129a83b301f2b2c0dbb606a5
	// 		Command being timed: "./sha-nostream"
	// 		User time (seconds): 0.26
	// 		System time (seconds): 0.07
	// 		Percent of CPU this job got: 121%
	// 		Elapsed (wall clock) time (h:mm:ss or m:ss): 0:00.28
	// 		Average shared text size (kbytes): 0
	// 		Average unshared data size (kbytes): 0
	// 		Average stack size (kbytes): 0
	// 		Average total size (kbytes): 0
	// 		Maximum resident set size (kbytes): 213636
	// 		Average resident set size (kbytes): 0
	// 		Major (requiring I/O) page faults: 0
	// 		Minor (reclaiming a frame) page faults: 53139
	// 		Voluntary context switches: 807
	// 		Involuntary context switches: 1
	//  	Swaps: 0
	// 		File system inputs: 0
	// 		File system outputs: 0
	// 		Socket messages sent: 0
	// 		Socket messages received: 0
	// 		Signals delivered: 0
	// 		Page size (bytes): 4096
	// 		Exit status: 0
	//
	//  But if we stream, we use less than 2M of memory:
	//
	// 		$ cat input.bin  | time -v ./sha-stream
	// 		e017fa5f3b78104b18c9e3ec00a678e513095cd4129a83b301f2b2c0dbb606a5
	// 		Command being timed: "./sha-stream"
	// 		User time (seconds): 0.16
	// 		System time (seconds): 0.02
	// 		Percent of CPU this job got: 99%
	// 		Elapsed (wall clock) time (h:mm:ss or m:ss): 0:00.18
	// 		Average shared text size (kbytes): 0
	// 		Average unshared data size (kbytes): 0
	// 		Average stack size (kbytes): 0
	// 		Average total size (kbytes): 0
	// 		Maximum resident set size (kbytes): 1948
	// 		Average resident set size (kbytes): 0
	// 		Major (requiring I/O) page faults: 0
	// 		Minor (reclaiming a frame) page faults: 198
	// 		Voluntary context switches: 89
	// 		Involuntary context switches: 19
	//  	Swaps: 0
	// 		File system inputs: 0
	// 		File system outputs: 0
	// 		Socket messages sent: 0
	// 		Socket messages received: 0
	// 		Signals delivered: 0
	// 		Page size (bytes): 4096
	// 		Exit status: 0
}
