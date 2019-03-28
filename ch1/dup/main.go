// dup counts duplicate lines either from stdin or list of arguments
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) <= 1 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, err = fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				if err != nil {
					panic(err)
				}
				continue
			}
			countLines(f, counts)
			_ = f.Close()
		}
	}

	for line, n := range counts {
		if n >= 2 {
			fmt.Printf("%d:\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	s := bufio.NewScanner(f)
	for s.Scan() {
		counts[s.Text()]++
	}
	if err := s.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "dup: %v\n", err)
	}
}
