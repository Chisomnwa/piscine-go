package piscine

func IterativePower(nb int, power int) int {
	// if power is negative
	if power < 0 {
		return 0
	}

	// if power equals 0
	if power == 0 {
		return 1
	}

	result := 1
	for i := 0; i < power; i++ {
		result *= nb
	}

	return result
}
