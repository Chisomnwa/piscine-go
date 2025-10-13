package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' { // if not a digit, return 0
			return 0
		}
		result = result*10 + int(ch-'0')
	}
	return result
}
