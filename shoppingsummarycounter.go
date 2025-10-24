package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for i := 0; i < len(str); i++ {
		char := str[i]

		if char != ' ' {
			word += string(char)
		} else {
			// Only count the word if it's not empty
			if word != "" {
				result[word]++
				word = ""
			}
		}
	}

	// Handle the last word (if there’s no trailing space)
	if word != "" {
		result[word]++
	}

	return result
}
