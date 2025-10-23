package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for a := 99; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			// Let's print the first number
			z01.PrintRune(rune(a/10 + '0'))
			z01.PrintRune(rune(a%10 + '0'))
			z01.PrintRune(' ')
			// Let's print the second number
			z01.PrintRune(rune(b/10 + '0'))
			z01.PrintRune(rune(b%10 + '0'))

			// Add comma + space, except for the last combination
			if !(a == 1 && b == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}

}
