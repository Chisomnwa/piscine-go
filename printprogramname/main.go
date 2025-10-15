package main

import ("os"
		"github.com/01-edu/z01")

func main() {
	// Get the program name (e.g., "./main" or "./printprogramme")
	programPath := os.Args[0]

	// Find only the part after the last '/'
	programName:= ""
	for i := len(programPath) - 1; i >= 0; i-- {
		if programPath[i] == '/' {
			programName = programPath[i+1:]
			break
		}
	} 

	// If there's no '/', the whole string is the program name
	if programName == "" {
		programName = programPath
	}

	// Print it rune by rune
	for _, r := range programName {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
