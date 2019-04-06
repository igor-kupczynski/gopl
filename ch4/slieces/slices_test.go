package slieces

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{"empty", []int{}, []int{}},
		{"1 elem", []int{1}, []int{1}},
		{"2 elems", []int{1, 2}, []int{2, 1}},
		{"3 elems", []int{1, 2, 3}, []int{3, 2, 1}},
		{"4 elems", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"10 elems", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := make([]int, len(tt.in))
			copy(op, tt.in)
			reverse(op)
			if !reflect.DeepEqual(tt.out, op) {
				t.Fatalf("%#v != %#v\n", tt.out, op)
			}
		})
	}
}

func Test_reverse2(t *testing.T) {
	in := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	expected := [10]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	reverse2(&in)
	if !reflect.DeepEqual(expected, in) {
		t.Fatalf("%#v != %#v\n", expected, in)
	}
}

func Test_rotateLeft(t *testing.T) {
	tests := []struct {
		name  string
		in    []int
		shift int
		out   []int
	}{
		{
			name:  "example",
			in:    []int{0, 1, 2, 3, 4, 5},
			shift: 2,
			out:   []int{2, 3, 4, 5, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := make([]int, len(tt.in))
			copy(op, tt.in)
			rotateLeft(op, tt.shift)
			if !reflect.DeepEqual(tt.out, op) {
				t.Fatalf("%#v != %#v\n", tt.out, op)
			}
		})
	}
}

func Test_uniq(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		out  []string
	}{
		{
			name: "empty",
			in:   []string{},
			out:  []string{},
		},
		{
			name: "short",
			in:   []string{"a"},
			out:  []string{"a"},
		},
		{
			name: "example",
			in:   []string{"foo", "foo", "bar", "baz", "baz", "baz", "foo"},
			out:  []string{"foo", "bar", "baz", "foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := make([]string, len(tt.in))
			copy(op, tt.in)
			op = uniq(op)
			if !reflect.DeepEqual(tt.out, op) {
				t.Fatalf("%#v != %#v\n", tt.out, op)
			}
		})
	}
}

func Test_squashSpace(t *testing.T) {
	tests := []struct {
		name string
		in   []rune
		out  []rune
	}{
		{
			name: "empty",
			in:   []rune{},
			out:  []rune{},
		},
		{
			name: "short",
			in:   []rune{'a'},
			out:  []rune{'a'},
		},
		{
			name: "example",
			in:   []rune{' ', ' ', '\t', 'f', ' ', ' ', 'f', ' ', '\t'},
			out:  []rune{' ', 'f', ' ', 'f', ' '},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := make([]rune, len(tt.in))
			copy(op, tt.in)
			op = squashSpaces(op)
			if !reflect.DeepEqual(tt.out, op) {
				t.Fatalf("%#v != %#v\n", tt.out, op)
			}
		})
	}
}
