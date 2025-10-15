package piscine

func IsAlpha(s string) bool {
	for _, char := range s {
		// check if not a letter or a digit
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9')) {
			return false
		}
	}
	return true
}
