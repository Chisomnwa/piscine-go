package piscine

func StrRev(s string) string {
	runes := []rune(s) // convert string to rune slice to handle Unicode correctly
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // swap runes
	}
	return string(runes) // convert back to string
}
