package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a { // loop through each string in the slice
		for _, char := range word { // loop through each character (rune)
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
