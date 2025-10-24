package piscine

func LoafOfBread(str string) string {
	// Step 1: remove spaces
	runes := []rune{}
	for _, r := range str {
		if r != ' ' {
			runes = append(runes, r)
		}
	}

	// Step 2: check if less than 5 chars
	if len(runes) < 5 {
		return "Invalid Output\n"
	}

	// Step 3: build result
	result := ""
	i := 0
	for {
		// Stop if less than 5 remain
		if i+5 > len(runes) {
			break
		}

		// Take 5 characters
		result += string(runes[i : i+5])
		i += 5

		// Skip one character *if there’s one to skip*
		if i < len(runes) {
			i++
		}

		// Add space if there are still enough letters left for another group
		if i+5 <= len(runes) {
			result += " "
		}
	}

	return result + "\n"
}
