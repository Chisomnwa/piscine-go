package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// If no arguments or help flag → show help text
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	insertStr := ""
	order := false
	mainStr := ""

	// Loop through arguments to detect flags
	for _, arg := range args {
		// Check for insert flag
		if len(arg) > 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			mainStr = arg // main string argument
		}
	}

	// Combine main string and insert string
	finalStr := mainStr + insertStr

	// If order flag is present, sort the string (ASCII order)
	if order {
		runes := []rune(finalStr)
		for i := 0; i < len(runes); i++ {
			for j := i + 1; j < len(runes); j++ {
				if runes[i] > runes[j] {
					runes[i], runes[j] = runes[j], runes[i]
				}
			}
		}
		finalStr = string(runes)
	}

	// Print final result
	for _, r := range finalStr {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// Prints the help instructions
func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}
