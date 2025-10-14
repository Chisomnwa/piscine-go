package piscine

func RecursivePower(nb int, power int) int {
	// iIF power is less than zero
	if power < 0 {
		return 0
	}

	// When power equals zero because anything raised 
	// to power zero is 1
	if power == 0 {
		return 1
	}

	// Calling the recursive
	return nb * RecursivePower(nb, power-1)
}
