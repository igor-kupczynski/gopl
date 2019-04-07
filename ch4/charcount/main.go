package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

type category byte

const (
	letter category = iota
	number
	punctuation
	space
)

func main() {

	var (
		// Count distinct runes
		chars = make(map[rune]int)

		// Rune legnth histogram
		lens [utf8.UTFMax]int

		// Numer of invalid characters
		invalid int

		// Count categories
		categories = make(map[category]int)
	)

	reader := bufio.NewReader(os.Stdin)
	for {
		r, n, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		chars[r]++
		lens[n]++

		if unicode.IsLetter(r) {
			categories[letter]++
		}
		if unicode.IsNumber(r) {
			categories[number]++
		}
		if unicode.IsPunct(r) {
			categories[punctuation]++
		}
		if unicode.IsSpace(r) {
			categories[space]++
		}

	}

	fmt.Printf("Rune\tCount\n")
	for r, n := range chars {
		fmt.Printf("%q\t%d\n", r, n)
	}

	fmt.Printf("\nRune length\n\tCount\n")
	for i, n := range lens {
		fmt.Printf("%d\t%d\n", i, n)
	}

	if invalid > 0 {
		fmt.Printf("\nInvalid runes: %d\n", invalid)
	}

	fmt.Printf("\nRune categories\tCount\n")
	fmt.Printf("Letters\t\t%d\n", categories[letter])
	fmt.Printf("Numbers\t\t%d\n", categories[number])
	fmt.Printf("Punctuations\t%d\n", categories[punctuation])
	fmt.Printf("Spaces\t\t%d\n", categories[space])

}
