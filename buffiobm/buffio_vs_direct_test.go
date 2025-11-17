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
			f.Write([]byte("x"))
		}
		f.Close()
	}
}

func BenchmarkBufioWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, _ := os.Create("bufio.txt")
		w := bufio.NewWriter(f)
		for j := 0; j < iterations; j++ {
			w.Write([]byte("x"))
		}
		w.Flush()
		f.Close()
	}
}
