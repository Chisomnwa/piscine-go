package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i, str := range strs {
		result += str
		if i < len(strs)-1 { // Add separator only between elements
			result += sep
		}
	}
	return result
}
