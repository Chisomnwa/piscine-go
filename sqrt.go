package piscine

func Sqrt(nb int) int {
	// Negative numbers do not have real square roots
	if nb < 0 {
		return 0
	}

	// Try possible integers starting from 1
	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}

	// If no integer square root found, return 0
	return 0
}
