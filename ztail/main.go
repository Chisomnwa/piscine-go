package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 || args[0] != "-c" {
		fmt.Println("Usage: go run . -c <number> <file> [<file>...]")
		os.Exit(1)
	}

	// Convert string to int
	count, err := strconv.Atoi(args[1])
	if err != nil || count < 0 {
		fmt.Println("Invalid count")
		os.Exit(1)
	}

	files := args[2:]
	exitCode := 0

	for i, file := range files {
		// Print header if multiple files
		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", file)
		}

		// Try reading file
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			exitCode = 1
			continue
		}

		// Calculate start index
		start := len(data) - count
		if start < 0 {
			start = 0
		}

		fmt.Print(string(data[start:]))
	}

	if exitCode != 0 {
		os.Exit(1)
	}
}
