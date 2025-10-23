package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args
	if len(a) < 4 || a[1] != "-c" {
		os.Exit(1)
	}
	n := 0
	for _, ch := range a[2] {
		if ch < '0' || ch > '9' {
			os.Exit(1)
		}
		n = n*10 + int(ch-'0')
	}
	hasError := false
	for i, f := range a[3:] {
		file, err := os.Open(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s: no such file or directory\n", f)
			hasError = true
			continue
		}
		info, _ := file.Stat()
		size := info.Size()
		bytesToRead := n
		if size < int64(n) {
			bytesToRead = int(size)
		}
		file.Seek(-int64(bytesToRead), 2)
		if len(a[3:]) > 1 {
			if i > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", f)
		}
		buf := make([]byte, bytesToRead)
		file.Read(buf)
		fmt.Printf("%s", buf)
		file.Close()
	}
	if hasError {
		os.Exit(1)
	}
}
