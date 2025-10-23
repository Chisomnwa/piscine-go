package piscine

func MyFunc(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		// Slices with 0, 1 or 2 element are always sorted
		return true
	}

	direction := 0 // 0 = unknown, -1 = ascending, 1 = descending
	for i := 0; i < len(a)-1 && direction == 0; i++ {
		direction = f(a[i], a[i+1])
	}

	if direction > 0 { // Direction is descending
		for i := 0; i < len(a)-1; i++ {
			// If we find an ascending pair, it's not sorted
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	} else { // Direction is ascending (or all elements are equal)
		for i := 0; i < len(a)-1; i++ {
			// If we find a descending pair, it's not sorted
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
	}

	return true
}
