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

func TestPopCountByteSlice(t *testing.T) {
	for _, tt := range tests {
		slice := make([]byte, 8)
		slice[0] = byte(tt.in >> (0 * 8))
		slice[1] = byte(tt.in >> (1 * 8))
		slice[2] = byte(tt.in >> (2 * 8))
		slice[3] = byte(tt.in >> (3 * 8))
		slice[4] = byte(tt.in >> (4 * 8))
		slice[5] = byte(tt.in >> (5 * 8))
		slice[6] = byte(tt.in >> (6 * 8))
		slice[7] = byte(tt.in >> (7 * 8))
		if got := PopCountByteSlice(slice); got != uint64(tt.want) {
			t.Errorf("testPopCountByteSlice() = %v, want %v", got, tt.want)
		}
	}
}

func TestPopCount(t *testing.T) {
	for _, tt := range tests {
		if got := PopCount(tt.in); got != tt.want {
			t.Errorf("popCount() = %v, want %v", got, tt.want)
		}
	}
}

func TestPopCountLookupTable(t *testing.T) {
	for _, tt := range tests {
		if got := popCountLookupTableExpr(tt.in); got != tt.want {
			t.Errorf("popCountLookupTableExpr() = %v, want %v", got, tt.want)
		}
	}
}

func TestPopCountLookupTableLoop(t *testing.T) {
	for _, tt := range tests {
		if got := popCountLookupTableLoop(tt.in); got != tt.want {
			t.Errorf("popCountLookupTableLoop() = %v, want %v", got, tt.want)
		}
	}
}

func TestPopCountNaive(t *testing.T) {
	for _, tt := range tests {
		if got := popCountNaive(tt.in); got != tt.want {
			t.Errorf("TestPopCountNaive() = %v, want %v", got, tt.want)
		}
	}
}

func TestPopCountClever(t *testing.T) {
	for _, tt := range tests {
		if got := popCountClever(tt.in); got != tt.want {
			t.Errorf("popCountClever() = %v, want %v", got, tt.want)
		}
	}
}

var (
	result       byte
	mostlyZeros  uint64 = 0x0001000100010001
	mostlyOnes   uint64 = 0xfffefffefffefffe
	zerosAndOnes uint64 = 0x0001fffe0001fffe
)

func benchmarkPopCount(b *testing.B, in uint64, popcount func(uint64) byte) {
	var r byte
	for i := 0; i < b.N; i++ {
		r = popcount(in)
	}
	result = r
}

func BenchmarkPopCountLookupTableMostlyZeros(b *testing.B) {
	benchmarkPopCount(b, mostlyZeros, popCountLookupTableExpr)
}
func BenchmarkPopCountLookupTableMostlyOnes(b *testing.B) {
	benchmarkPopCount(b, mostlyOnes, popCountLookupTableExpr)
}
func BenchmarkPopCountLookupTableZerosAndOnes(b *testing.B) {
	benchmarkPopCount(b, zerosAndOnes, popCountLookupTableExpr)
}

func BenchmarkPopCountLookupTableLoopMostlyZeros(b *testing.B) {
	benchmarkPopCount(b, mostlyZeros, popCountLookupTableLoop)
}
func BenchmarkPopCountLookupTableLoopMostlyOnes(b *testing.B) {
	benchmarkPopCount(b, mostlyOnes, popCountLookupTableLoop)
}
func BenchmarkPopCountLookupTableLoopZerosAndOnes(b *testing.B) {
	benchmarkPopCount(b, zerosAndOnes, popCountLookupTableLoop)
}

func BenchmarkPopCountNaiveMostlyZeros(b *testing.B) {
	benchmarkPopCount(b, mostlyZeros, popCountNaive)
}
func BenchmarkPopCountNaiveMostlyOnes(b *testing.B) {
	benchmarkPopCount(b, mostlyOnes, popCountNaive)
}
func BenchmarkPopCountNaiveZerosAndOnes(b *testing.B) {
	benchmarkPopCount(b, zerosAndOnes, popCountNaive)
}

func BenchmarkPopCountCleverMostlyZeros(b *testing.B) {
	benchmarkPopCount(b, mostlyZeros, popCountClever)
}
func BenchmarkPopCountCleverMostlyOnes(b *testing.B) {
	benchmarkPopCount(b, mostlyOnes, popCountClever)
}
func BenchmarkPopCountCleverZerosAndOnes(b *testing.B) {
	benchmarkPopCount(b, zerosAndOnes, popCountClever)
}

// Benchmark results
//
// BenchmarkPopCountLookupTableMostlyZeros-8        	500000000	         3.38 ns/op
// BenchmarkPopCountLookupTableMostlyOnes-8         	500000000	         3.33 ns/op
// BenchmarkPopCountLookupTableZerosAndOnes-8       	500000000	         3.38 ns/op
// BenchmarkPopCountLookupTableLoopMostlyZeros-8    	100000000	        14.1 ns/op
// BenchmarkPopCountLookupTableLoopMostlyOnes-8     	100000000	        14.3 ns/op
// BenchmarkPopCountLookupTableLoopZerosAndOnes-8   	100000000	        14.0 ns/op
// BenchmarkPopCountNaiveMostlyZeros-8              	30000000	        41.1 ns/op
// BenchmarkPopCountNaiveMostlyOnes-8               	30000000	        41.3 ns/op
// BenchmarkPopCountNaiveZerosAndOnes-8             	30000000	        41.1 ns/op
// BenchmarkPopCountCleverMostlyZeros-8             	500000000	         3.05 ns/op
// BenchmarkPopCountCleverMostlyOnes-8              	50000000	        27.2 ns/op
// BenchmarkPopCountCleverZerosAndOnes-8            	100000000	        13.7 ns/op
//
// We have 4 pop count methods:
// (1) lookup table,
// (2) lookup table with a loop,
// (3) naive -- shift and count every bit
// (4) clever -- x&(x-1) zeros the least significant "1", so we can count the
//               "1s" this wahy
//
//
// Except for (4), the other methods are stable, that is the input doesn't matter,
// the cost is the same.
//
// Unsurprisingly, the naive methods is the slowest (more than 10x slower than
// the fastest). The clever one is the fastest, but only for inputs
// with mostly "0s" in their binary representation.
//
// The lookup table does really good, being only slightly slower than clever
// in clever's best cast and significantly faster in its avg and worst case.
//
// It makes sense to manually unroll the loop, as the perf gain is 4x.
