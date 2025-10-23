package piscine

func CountIf(f func(string) bool, tab []string) int {
	var counter int = 0

	for _, val := range tab {
		if f(val) {
			counter++
		}
	}
	return counter
}
