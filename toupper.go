package piscine

func ToUpper(s string) string {
	runes := []rune(s) // convert string to runes (handles Unicode safely)
	for i, char := range runes {
		if char >= 'a' && char <= 'z' {
			// difference between lowercase and uppercase in ASCII is 32
			runes[i] = char - 32
		}
	}
	return string(runes)
}
