package slieces

import "unicode"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// reverse2 is same as reverse but uses an array pointer.
func reverse2(s *[10]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// rotateLeft rotates slice s by shift positions left.
func rotateLeft(s []int, shift int) {
	n := len(s)

	residue := make([]int, len(s[:shift]))
	copy(residue, s[:shift])

	copy(s[:n-shift], s[shift:])
	copy(s[n-shift:], residue)
}

// uniq removes adjacent duplicates.
// The underlying slice is modified during the call.
func uniq(s []string) []string {
	if len(s) < 2 {
		return s
	}
	last := s[0]
	writer := 1
	for reader := 1; reader < len(s); reader++ {
		if s[reader] != last {
			s[writer] = s[reader]
			writer++
		}
		last = s[reader]
	}
	return s[:writer]
}

// squashSpaces squashes all unicode spaces of an utf-8 encoded string.
// The underlying slice is modified during the call.
func squashSpaces(s []rune) []rune {
	if len(s) < 2 {
		return s
	}
	last := ' '
	writer := 1
	for reader := 1; reader < len(s); reader++ {
		if unicode.IsSpace(s[reader]) {
			if last != ' ' {
				s[writer] = ' '
				writer++
			}
		} else {
			s[writer] = s[reader]
			writer++
		}
		last = s[reader]
	}
	return s[:writer]
}
