package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert nbr (string) in baseFrom → decimal (int)
	baseFromLen := len(baseFrom)
	baseToLen := len(baseTo)

	// Step 1: baseFrom → base 10
	decimalValue := 0
	for _, digit := range nbr {
		index := 0
		for i, c := range baseFrom {
			if c == digit {
				index = i
				break
			}
		}
		decimalValue = decimalValue*baseFromLen + index
	}

	// Step 2: base 10 → baseTo
	if decimalValue == 0 {
		return string(baseTo[0])
	}

	result := ""
	for decimalValue > 0 {
		remainder := decimalValue % baseToLen
		result = string(baseTo[remainder]) + result
		decimalValue /= baseToLen
	}

	return result
}
