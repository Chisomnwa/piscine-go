package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	isNewWord := true

	for i, r := range runes {
		// Check if alphanumeric (A–Z, a–z, 0–9)
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if isNewWord {
				// Capitalize first letter if it's a lowercase
				if r >= 'a' && r <= 'z' {
					runes[i] = r - 32
				}
				isNewWord = false
			} else {
				// Lowercase any uppercase letters after the first
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + 32
				}
			}
		} else {
			// Any non-alphanumeric means next char starts a new word
			isNewWord = true
		}
	}
	return string(runes)
}
