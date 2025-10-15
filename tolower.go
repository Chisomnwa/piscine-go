package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			// difference between uppercase and lowercase in ASCII is 32
			runes[i] = char + 32
		}
	}
	return string(runes)
}
