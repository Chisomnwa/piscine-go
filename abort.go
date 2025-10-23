package piscine

func Abort(a, b, c, d, e int) int {
	num := []int{a, b, c, d, e}

	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num)-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	return num[2]
}
