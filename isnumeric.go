package piscine

func IsNumeric(s string) bool {
	if s == "" {
		return false
	}

	for _, number := range s {
		if number < '0' || number > '9' {
			return false
		}
	}
	return true
}
