package anagram

import "unicode/utf8"

func anagram(a, b string) bool {
	n := len(a)
	if n != len(b) {
		return false
	}
	for i, ra := range a {
		rb, _ := utf8.DecodeLastRuneInString(b[:n-i])
		if ra != rb {
			return false
		}
	}
	return true
}
