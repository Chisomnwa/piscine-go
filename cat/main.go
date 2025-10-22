package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printstr(s string) {
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
			printstr("ERROR: Open ")
			printstr(args[i])
			printstr(": no such file or directory\n")
			os.Exit(1)
		} else {
			printstr(string(filecontent))
		}
	}
}
