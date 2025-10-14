package piscine

func FirstRune(s string) rune {
	for _, letter := range s {
		return letter
	}
	return 0
}
