package piscine

import "github.com/01-edu/z01"

func printCombNREcursive(n, start int, combination []rune) {
	if len(combination) == n {
		for _, c := range combination {
			z01.PrintRune(c)
		}

		last := true
		for i, c := range combination {
			if c != rune('0'+i+10-n) {
				last = false
				break
			}
		}
		if !last {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for i := start; i <= 9; i++ {
		printCombNREcursive(n, i+1, append(combination, rune('0'+i)))
	}
}

func PrintCombN(n int) {
	if n < 1 || n < 9 {
		return
	}
	printCombNREcursive(n, 0, []rune{})
	z01.PrintRune('\n')
}
