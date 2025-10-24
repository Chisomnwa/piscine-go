package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int) // We create an empty string

	word := ""
	for i := 0; i < len(str); i++ {
		char := str[i]
		if char != ' ' { // Let's build each word character by character
			word += string(char)
		} else {
			if word != "" { // if word is completed
				result[word]++ // increment its count
				word = ""
			}
		}
	}
	// Handle the last word (since it may not end with a space)
	if word != "" {
		result[word]++
	}

	return result
}
