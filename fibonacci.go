package piscine

func Fibonacci(index int) int {
	// Negative index should return -1
	if index < 0 {
		return -1
	}

	// Base cases
	if index == 0 {
		return 0
	}

	if index == 1 {
		return 1
	}

	// Now calling the recursive case
	return Fibonacci(index-1) + Fibonacci(index-2)
}
