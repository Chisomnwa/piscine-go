package piscine

func Fibonacci(index int) int {
	// NEgative index should return -1
	if index < 0 {
		return -1
	}

	// Base cases: Fib(0) = 0 and Fib(1) = 1
	if index == 0 {
		return 0
	}
	
	if index == 1 {
		return 1
	}

	// NOw calling the recusrive
	return Fibonacci(index-1) + Fibonacci(index-2)
	}
