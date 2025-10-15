package piscine

func IsPrintable(s string) bool {
	for _, character := range s {
		if character < 32 || character > 126 {
			return false
		}
	}
	return true
}
