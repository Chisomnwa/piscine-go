package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	// Step 1: Check if base is valid
	if !isValidBase(base) {
		for _, r := range "NV" {
			z01.PrintRune(r)
		}
		return
	}

	// Step 2: Handle negative numbers
	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 { // handle int64 min overflow
			printNumber(-(nbr / len(base)), base)
			z01.PrintRune(rune(base[-(nbr % len(base))]))
			return
		}
		nbr = -nbr
	}

	// Step 3: Print the number in the given base
	printNumber(nbr, base)
}

// Helper to print the digits recursively
func printNumber(n int, base string) {
	baseLen := len(base)
	if n >= baseLen {
		printNumber(n/baseLen, base)
	}
	z01.PrintRune(rune(base[n%baseLen]))
}

// Helper to check if the base is valid
func isValidBaseAtoi(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i, ch := range base {
		if ch == '+' || ch == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if ch == rune(base[j]) {
				return false // Duplicate found
			}
		}
	}
	return true
}
