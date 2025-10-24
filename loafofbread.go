package piscine

func LoafOfBread(str string) string {
	// Step 1: Remove spaces from input
	runes := []rune{}
	for _, r := range str {
		if r != ' ' {
			runes = append(runes, r)
		}
	}

	// Step 2: Check minimum length
	if len(runes) < 5 {
		return "Invalid Output\n"
	}

	// Step 3: Build result
	result := ""
	i := 0
	for i+5 <= len(runes) {
		// Take 5 characters
		result += string(runes[i : i+5])

		i += 5 // move past 5 characters

		// Add a space if there are more characters after skipping
		if i < len(runes) {
			result += " "
		}

		i++ // skip 1 character
	}

	return result + "\n"
}
