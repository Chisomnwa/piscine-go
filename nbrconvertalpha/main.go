package main

import (
	"os"

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

	printed := false // Tracks if we printed anything at all

	for _, arg := range args {
		num, ok := toInt(arg) // convert string to integer
		if !ok || num < 1 || num > 26 {
			// Invalid argument - prints space
			z01.PrintRune(' ')
			printed = true
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
		printed = true
	}

	// Only print newline if something was printed
	if printed {
		z01.PrintRune('\n')
	}
}

// Converts a numeric string to an integer manually (no strconv)
func toInt(s string) (int, bool) {
	num := 0
	if s == "" { // empty string is invalid
		return 0, false
	}
	for _, r := range s {
		if r < '0' || r > '9' { // only digits allowed
			return 0, false
		}
		num = num*10 + int(r-'0') // build integer manually
	}
	return num, true
}
