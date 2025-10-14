package piscine

func RecursiveFactorial(nb int) int {
	// Because negative numbers do not have factorials 
	if nb < 0 {
		return 0
	}

	// Base case: 0! = 1 and 1! = 1
	if nb == 0 || nb == 1 {
		return 1
	}

	// Recursive call
	result := nb * RecursiveFactorial(nb -1)

	// Check for overflow (If result becomes negative)
	if result < 0 {
		return 0
	}

	return result
}
