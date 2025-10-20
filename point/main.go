package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{} // CReate a ponter to a new 'point' struct
	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
