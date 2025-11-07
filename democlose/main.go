package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	dir, err := ioutil.TempDir("", "leaktest")
	if err != nil {
		panic(err)
	}
	fmt.Println("temp dir:", dir)

	const N = 100000
	for i := 0; i < N; i++ {
		name := filepath.Join(dir, fmt.Sprintf("tmp-%d", i))
		// create file (so we have many distinct file entries)
		if err := ioutil.WriteFile(name, []byte("x"), 0644); err != nil {
			fmt.Printf("create error at %d: %v\n", i, err)
			return
		}

		// open file and DO NOT close it
		f, err := os.Open(name)
		if err != nil {
			fmt.Printf("open error at %d: %v\n", i, err)
			return
		}

		// keep `f` alive (not closed)
		_ = f

		if i%100 == 0 {
			fmt.Printf("opened %d files\n", i)
		}
	}
}
