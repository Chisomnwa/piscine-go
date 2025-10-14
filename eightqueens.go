package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var positions [8]int
	solve(0, &positions)
}

func solve(col int, positions *[8]int) {
	if col == 8 {
		printSolution(positions)
		return
	}

	for row := 1; row <= 8; row++ {
		if isSafe(col, row, positions) {
			positions[col] = row
			solve(col+1, positions)
		}
	}
}

func isSafe(col, row int, positions *[8]int) bool {
	for prevCol := 0; prevCol < col; prevCol++ {
		prevRow := positions[prevCol]
		// Check same row or diagonal
		if prevRow == row || abs(prevRow-row) == abs(prevCol-col) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printSolution(positions *[8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(positions[i]) + '0')
	}
	z01.PrintRune('\n')
}
