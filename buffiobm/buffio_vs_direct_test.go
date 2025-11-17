package main

import (
	"bufio"
	"os"
	"testing"
)

const iterations = 10000

func BenchmarkDirectWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, _ := os.Create("direct.txt")
		for j := 0; j < iterations; j++ {
			_, _ = f.Write([]byte("x"))
		}
		f.Close()
	}
}

func BenchmarkBufioWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, _ := os.Create("bufio.txt")
		w := bufio.NewWriter(f)
		for j := 0; j < iterations; j++ {
			_, _ = w.Write([]byte("x"))
		}
		w.Flush()
		f.Close()
	}
}

/*
go test -bench=.

goos: linux
goarch: amd64
pkg: buffiobm
cpu: Intel(R) Core(TM) i5-10310U CPU @ 1.70GHz
BenchmarkDirectWrite-8               133           9157738 ns/op
BenchmarkBufioWrite-8               5072            211429 ns/op
PASS
ok      buffiobm        3.225s

*/
