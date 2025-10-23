package piscine

func Unmatch(a []int) int {
	count := 0
	for range a {
		count++
	}

	for i := 0; i < count; i++ {
		repeat := 0
		for j := 0; j < count; j++ {
			if a[i] == a[j] {
				repeat++
			}
		}
		if repeat%2 != 0 {
			return a[i]
		}
	}
	return -1
}
