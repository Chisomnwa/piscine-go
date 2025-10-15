package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// Sort the args in ASCII order (simple bubble sort)
	for i := 0; i < len(args)-1; i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	// PRinted sorted arguments
	for _, word := range args {
		for _, r := range word {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
