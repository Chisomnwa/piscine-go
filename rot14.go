package piscine

func Rot14(s string) string {
	var result string
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r = ((r - 'a' + 14) % 26) + 'a'
		} else if r >= 'A' && r <= 'Z' {
			r = ((r - 'A' + 14) % 26) + 'A'
		}
		result += string(r)
	}
	return result
}
