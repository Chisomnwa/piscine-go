package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, ch := range s {
		// Convert rune '0'...'9' to int by subtracting '0' rune
		result = result*10 + int(ch-'0')
	}
	return result
}
