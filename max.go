package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	// Assume the first element is the max
	max := a[0]

	// Loop through the slice manually
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	return max
}
