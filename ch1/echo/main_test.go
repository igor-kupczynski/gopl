package main

import "testing"

var result string
var (
	args10 = generateArgs(10)
	args100 = generateArgs(100)
	args1000 = generateArgs(1000)
	args10000 = generateArgs(10000)
	args100000 = generateArgs(100000)
)

func generateArgs(n int) []string {
	seed := []string { "razzmatazz", "bemuzzling", "puzzlingly", "whizzbangs",
		"embezzling", "unmuzzling", "unpuzzling", "blackjacks", "dizzyingly",
		"puzzlement", "scuzzballs", "zigzagging", "bedazzling",}
	args := make([]string, n)
	for i := 0; i < n; i++ {
		args[i] = seed[i % len(seed)]
	}
	return args
}

func benchmarkNaiveEcho(b *testing.B, args []string) {
	var r string
	for i := 0; i < b.N; i++ {
		r = naiveEcho(args)
	}
	result = r
}

func benchmarkLibEcho(b *testing.B, args []string) {
	var r string
	for i := 0; i < b.N; i++ {
		r = libEcho(args)
	}
	result = r
}

func BenchmarkNaiveEcho10(b *testing.B) { benchmarkNaiveEcho(b, args10) }
func BenchmarkNaiveEcho100(b *testing.B) { benchmarkNaiveEcho(b, args100) }
func BenchmarkNaiveEcho1000(b *testing.B) { benchmarkNaiveEcho(b, args1000) }
func BenchmarkNaiveEcho10000(b *testing.B) { benchmarkNaiveEcho(b, args10000) }
func BenchmarkNaiveEcho100000(b *testing.B) { benchmarkNaiveEcho(b, args100000) }

func BenchmarkLibEcho10(b *testing.B) { benchmarkLibEcho(b, args10) }
func BenchmarkLibEcho100(b *testing.B) { benchmarkLibEcho(b, args100) }
func BenchmarkLibEcho1000(b *testing.B) { benchmarkLibEcho(b, args1000) }
func BenchmarkLibEcho10000(b *testing.B) { benchmarkLibEcho(b, args10000) }
func BenchmarkLibEcho100000(b *testing.B) { benchmarkLibEcho(b, args100000) }

