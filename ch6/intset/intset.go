package intset

import (
	"bytes"
	"fmt"
)

// wsize represents the number of bits in uint
const wsize = 32 << (^uint(0) >> 63)

// IntSet is a set of small non-negative integers. Zero value is an empty set.
type IntSet struct {
	words []uint
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/wsize, uint(x%wsize)
	for len(s.words) <= word {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds all elements xs to the set.
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	t := &IntSet{}
	if s.words != nil {
		t.words = make([]uint, len(s.words))
		copy(t.words, s.words)
	}
	return t
}

// Elems return the elements present in the set as slice.
func (s *IntSet) Elems() []int {
	var elems []int
	for n, word := range s.words {
		for bit := uint(0); bit < wsize; bit++ {
			if word&(1<<bit) != 0 {
				elems = append(elems, wsize*n+int(bit))
			}
		}
	}
	return elems
}

// Has checks if the set contains non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit := x/wsize, uint(x%wsize)
	return len(s.words) > word && s.words[word]&(1<<bit) != 0
}

// Len returns the number of elements in the set
func (s *IntSet) Len() int {
	sum := 0
	for _, word := range s.words {
		for bit := uint(0); bit < wsize; bit++ {
			if word&(1<<bit) != 0 {
				sum++
			}
		}
	}
	return sum
}

// Remove removes x form the set
func (s *IntSet) Remove(x int) {
	word, bit := x/wsize, uint(x%wsize)
	if len(s.words) > word {
		s.words[word] &^= 1 << bit
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wsize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wsize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
