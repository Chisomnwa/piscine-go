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

func main() {
	points := &point{}
	setPoint(points)

	// Define slice as rune instead  of []int (so no conversion  needed)
	output := []rune{
		120, // x
		32,  // space
		61,  // =
		32,  // space
		52,  // '4'
		50,  // '2'
		44,  // '.'
		32,  // space
		121, // y
		32,  // space
		61,  // =
		50,  // '2'
		49,  // '1'
		10,  // newline
	}

	for _, v := range output {
		z01.PrintRune(v)
	}
}
