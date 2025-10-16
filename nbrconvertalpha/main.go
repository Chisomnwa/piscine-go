package main

import (
	"os"

	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	// Check for the upper flag when passed as one of the command line args
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:] // Remove the flag from arguments
	}

	for _, arg := range args {
		num, err := strconv.Atoi(arg) // convert string to integer
		if err != nil || num < 1 || num > 20 {
			// Invalid argument - prints space
			z01.PrintRune(' ')
			continue
		}

		// Convert number to corresponding letter
		var letter rune
		if upper {
			letter = rune('A' + num - 1) // uppercase letters
		} else {
			letter = rune('a' + num - 1) // lowercase letters
		}

		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
