package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, n := range a {
		result[i] = f(n)
	}
	return result
}
