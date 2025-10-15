package piscine

func NRune(s string, n int) rune {
	count := 0
	for _, letter := range s {
		count++
		if count == n && count <= len(s) {
			return letter
		}
	}
	return rune(0)
}
