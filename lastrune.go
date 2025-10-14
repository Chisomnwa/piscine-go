package piscine

func LastRune(s string) rune {
	runes := []rune(s)         // convert   strings to slice of bytes to habdle correctlly
	return runes[len(runes)-1] // Return the last rune
}
