package popcount

import (
	"testing"
)

var tests = []struct {
	in   uint64
	want byte
}{
	{0x0, 0},
	{0x1, 1},
	{0x2, 1},
	{0x3, 2},
	{0x4, 1},
	{0x5, 2},
	{0xfff0, 12},
	{0x1000000100000001, 3},
	{0xffff000000000000, 16},
	{0xfffefffffeffffef, 61},
	{0xffffffffffffffff, 64},
}

func TestPopCountLookupTable(t *testing.T) {
	for _, tt := range tests {
		if got := PopCountLookupTableExpr(tt.in); got != tt.want {
			t.Errorf("PopCountLookupTableExpr() = %v, want %v", got, tt.want)
		}
	}
}

func TestPopCountLookupTableLoop(t *testing.T) {
	for _, tt := range tests {
		if got := PopCountLookupTableLoop(tt.in); got != tt.want {
			t.Errorf("PopCountLookupTableLoop() = %v, want %v", got, tt.want)
		}
	}
}

func TestPopCountNaive(t *testing.T) {
	for _, tt := range tests {
		if got := PopCountNaive(tt.in); got != tt.want {
			t.Errorf("TestPopCountNaive() = %v, want %v", got, tt.want)
		}
	}
}

func TestPopCountClever(t *testing.T) {
	for _, tt := range tests {
		if got := PopCountClever(tt.in); got != tt.want {
			t.Errorf("PopCountClever() = %v, want %v", got, tt.want)
		}
	}
}
