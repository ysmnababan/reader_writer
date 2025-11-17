package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
)

func main() {
	printMem("init")
	size := 50 * 1024 * 1024
	data := bytes.Repeat([]byte("A"), size)
	src := bytes.NewReader(data)

	printMem("start")

	// ============ EXAMPLE 1: BUFFER EVERYTHING ====================
	buf, err := io.ReadAll(src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ReadAll finished: got %d bytes\n", len(buf))
	
	printMem("after ReadAll")

	// ====== EXAMPLE 2: STREAM DATA ======
	_, _ = src.Seek(0, 0)             // reset source
	dst, _ := os.Create("output.dat") // file on disk
	defer dst.Close()

	_, _ = io.Copy(dst, src)
	fmt.Println("Copy finished")

	printMem("afer io.Copy")
}

func printMem(stage string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s] Memory: %.2f MB\n", stage, float64(m.Alloc)/1024/1024)
}
