// Benchmarks for the two string concatenation approaches in gopl/ch1/echo
// I've read on them https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
//
// In general it looks that stdlib approach to string concat is 1000-10000 times
// than a naive approach for large (10k, 100k) inputs, and few times faster
// for small inputs (10, 100).
//
// Nevertheless, for small inputs probably most of the time will be spent in
// writing to terminal.
//
// Raw Results:
//
// GOROOT=/snap/go/current #gosetup
// GOPATH=/home/igor/code/golang #gosetup
// /snap/go/current/bin/go test -c -o /tmp/___gobench_github_com_igor_kupczynski_gopl_ch1_echo github.com/igor-kupczynski/gopl/ch1/echo #gosetup
// /tmp/___gobench_github_com_igor_kupczynski_gopl_ch1_echo -test.v -test.bench . -test.run ^$ -test.benchtime=60s #gosetup
//
// goos: linux
// goarch: amd64
// pkg: github.com/igor-kupczynski/gopl/ch1/echo
// BenchmarkNaiveEcho10-8       	100000000	       727 ns/op
// BenchmarkNaiveEcho100-8      	 3000000	     21828 ns/op
// BenchmarkNaiveEcho1000-8     	  100000	   1227328 ns/op
// BenchmarkNaiveEcho10000-8    	    1000	  83767983 ns/op
// BenchmarkNaiveEcho100000-8   	      10	7015283853 ns/op
// BenchmarkLibEcho10-8         	1000000000	       125 ns/op
// BenchmarkLibEcho100-8        	100000000	      1092 ns/op
// BenchmarkLibEcho1000-8       	10000000	     11239 ns/op
// BenchmarkLibEcho10000-8      	 1000000	    106953 ns/op
// BenchmarkLibEcho100000-8     	  100000	   1012886 ns/op
// PASS
package main

import "testing"

var result string
var (
	args10     = generateArgs(10)
	args100    = generateArgs(100)
	args1000   = generateArgs(1000)
	args10000  = generateArgs(10000)
	args100000 = generateArgs(100000)
)

func generateArgs(n int) []string {
	seed := []string{"razzmatazz", "bemuzzling", "puzzlingly", "whizzbangs",
		"embezzling", "unmuzzling", "unpuzzling", "blackjacks", "dizzyingly",
		"puzzlement", "scuzzballs", "zigzagging", "bedazzling"}
	args := make([]string, n)
	for i := 0; i < n; i++ {
		args[i] = seed[i%len(seed)]
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

func BenchmarkNaiveEcho10(b *testing.B)     { benchmarkNaiveEcho(b, args10) }
func BenchmarkNaiveEcho100(b *testing.B)    { benchmarkNaiveEcho(b, args100) }
func BenchmarkNaiveEcho1000(b *testing.B)   { benchmarkNaiveEcho(b, args1000) }
func BenchmarkNaiveEcho10000(b *testing.B)  { benchmarkNaiveEcho(b, args10000) }
func BenchmarkNaiveEcho100000(b *testing.B) { benchmarkNaiveEcho(b, args100000) }

func BenchmarkLibEcho10(b *testing.B)     { benchmarkLibEcho(b, args10) }
func BenchmarkLibEcho100(b *testing.B)    { benchmarkLibEcho(b, args100) }
func BenchmarkLibEcho1000(b *testing.B)   { benchmarkLibEcho(b, args1000) }
func BenchmarkLibEcho10000(b *testing.B)  { benchmarkLibEcho(b, args10000) }
func BenchmarkLibEcho100000(b *testing.B) { benchmarkLibEcho(b, args100000) }
