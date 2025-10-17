package piscine

func Split(s, sep string) []string {
	var result []string
	sepLen :=  len(sep)
	start := 0

	// iterate through string s
	for i := 0; i+sepLen <= len(s); i++ {
		// check if subtracting from i to i+sepLen matches sep
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i += sepLen - 1
			
		}
	}

	// append last part (after the last seperator)
	result = append(result, s[start:])

	return result
}
