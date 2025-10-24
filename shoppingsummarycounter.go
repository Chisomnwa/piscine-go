package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for i := 0; i < len(str); i++ {
		char := str[i]

		if char != ' ' {
			word += string(char)
		} else {
			// even if word is empty, count it
			result[word]++
			word = ""
		}
	}

	// count the last word (even if it's empty)
	result[word]++

	return result
}
