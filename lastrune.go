package piscine

// LastRune returns the last rune (Unicode code point) of a string.
// if the string is empty, it returns rune(0)
func LastRune(s string) rune {
	var last rune
	for _, letter := range s {
		last = letter
	}
	return last
}
