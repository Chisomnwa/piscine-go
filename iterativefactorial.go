package piscine

func IterativeFactorial(nb int) int {
	// This function returns the factorial of the parameter as an int
	if nb < 0 {
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ {
		result = result * i
	}
	return result
}
