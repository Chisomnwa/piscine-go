package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for x := '0' ; x <= '9' ; x++ {
		for y := '1' ; y < x ; y++ {
			for z := '2' ; z < y ; z++ {
				z01.PrintRune(z)
				z01.PrintRune(y)
				z01.PrintRune(x)
				if !(z == '7' && y == '8' && x == '9') {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
z01.PrintRune(rune('\n'))
}