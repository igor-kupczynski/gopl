package popcount

// pc[i] is the population (bit) count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountByteSlice(data []byte) uint64 {
	var sum uint64
	for _, b := range data {
		sum += uint64(pc[b])
	}
	return sum
}

func PopCount(x uint64) byte {
	return popCountLookupTableExpr(x)
}

func popCountLookupTableExpr(x uint64) byte {
	return pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))]
}

func popCountLookupTableLoop(x uint64) byte {
	var sum byte
	for i := byte(0); i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return sum
}

func popCountNaive(x uint64) byte {
	var sum byte
	for i := byte(0); i < 64; i++ {
		sum += byte((x >> i) & 1)
	}
	return sum
}

func popCountClever(x uint64) byte {
	var sum byte
	for x != 0 {
		sum++
		// clear the rightmost non-zero bit
		x = x & (x - 1)
	}
	return sum
}
