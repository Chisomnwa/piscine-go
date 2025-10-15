package piscine

func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		// Check if it's a lower case letter
		// i. e if the if the rune’s numeric value is between 97 and 122
		if char >= 'a' && char <= 'z' {
			count++
		}

		// Check if it's an upper case letter
		// i. e if the rune's numeric value is between 65 and 90
		if char >= 'A' && char <= 'Z' {
			count++
		}
	}

	return count
}
