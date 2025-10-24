package piscine

func LoafOfBread(str string) string {
	// Step 1: Remove spaces and keep only valid characters
	chars := []rune{}
	for _, r := range str {
		if r != ' ' {
			chars = append(chars, r)
		}
	}

	// Step 2: If less than 5 characters, invalid
	if len(chars) < 5 {
		return "Invalid Output\n"
	}

	// Step 3: Build result with 5 chars, skip 1 each time
	result := ""

	for i := 0; i < len(chars); {
		// Take 5 characters
		for j := 0; j < 5 && i < len(chars); j++ {
			result += string(chars[i])
			i++
		}

		// Add a space if there are more characters
		if i < len(chars) {
			result += " "
		}

		// Skip one character
		i++
	}

	// Step 4: Return result with newline
	return result + "\n"
}
