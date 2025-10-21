package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('-')
		n = -n
	}

	digits := []rune{}
	for n > 0 {
		digits = append([]rune{rune(n%10 + '0')}, digits...)
		n /= 10
	}

	for _, d := range digits {
		z01.PrintRune(d)
	}
}

func main() {
	points := &point{}
	setPoint(points)

	printStr("x =")
	printInt(points.x)
	printStr(", y =")
	printInt(points.y)
	z01.PrintRune('\n')
}
