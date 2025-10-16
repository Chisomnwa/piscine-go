package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Step 1: Collect all vowels
	var vowels []rune
	for _, arg := range args {C
		for _, r := range arg {
			if isVowel(r) {
				vowels = append(vowels, r)
			}
		}
	}

	// Step 2: If no vowels, print args as-is
	if len(vowels) == 0 {
		printArgs(args)
		return
	}

	// Step 3: Reverse vowels slice
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// Step 4: Replace vowels in original order with reversed vowels
	vowelIndex := 0
	for i, arg := range args {
		for _, r := range arg {
			if isVowel(r) {
				z01.PrintRune(vowels[vowelIndex])
				vowelIndex++
			} else {
				z01.PrintRune(r)
			}
		}
		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

// Checks if a rune is a vowel (case-insensitive)
func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

// Prints args separated by spaces
func printArgs(args []string) {
	for i, arg := range args {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
