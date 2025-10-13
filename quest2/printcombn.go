package piscine

import "github.com/01-edu/z01"

func printCombNRecursive(n, start int, combination []rune) {
	if len(combination) == n {
		for _, c := range combination {
			z01.PrintRune(c)
		}

		// Detect if this is the last combination
		if combination[0] != rune('0'+(10-n)) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for i := start; i <= 9; i++ {
		printCombNRecursive(n, i+1, append(combination, rune('0'+i)))
	}
}

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	printCombNRecursive(n, 0, []rune{})
	z01.PrintRune('\n')
}
