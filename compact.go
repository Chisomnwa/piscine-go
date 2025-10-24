package piscine

func Compact(ptr *[]string) int {
	s := *ptr
	count := 0

	// First, move all non-empty strings to the front manually
	for i := range s {
		if s[i] != " " {
			s[count] = s[i]
			count++
		}
	}

	// Then truncate the slice to contain only the non-empty elements
	*ptr = s[:count]
	return count
}
