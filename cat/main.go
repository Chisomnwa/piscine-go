package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printstrs(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for i := 0; i < len(args); i++ {
		filecontent, err := os.ReadFile(args[i])
		if err != nil {
			printstrs("ERROR: open ")
			printstrs(args[i])
			printstrs(": no such file or directory\nexit status 1\n")
		} else {
			printstrs(string(filecontent))
		}
	}
}
