package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	for w, n := range words {
		fmt.Printf("%d\t%s\n", n, w)
	}
}
