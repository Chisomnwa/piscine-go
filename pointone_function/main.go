package main

import (
	"fmt"

	"piscine"
)

// For calling the function PointOne from piscine-go/pointone.go
func main() {
	n := 0
	piscine.PointOne(&n)
	fmt.Println(n)
}
