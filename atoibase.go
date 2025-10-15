package piscine

func AtoiBase(s string, base string) int {
	// Step 1: Validate the base
	if !isValidBase(base) {
		return 0
	}

	// Step 2: Create a lookup map for base characters → their numeric values
	baseLen := len(base)
	valueMap := make(map[rune]int)
	for i, r := range base {
		valueMap[r] = i
	}

	// Step 3: Convert the string to an integer
	result := 0
	for _, r := range s {
		result = result*baseLen + valueMap[r]
	}

	return result
}

// Helper to check if base is valid
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i, ch := range base {
		if ch == '+' || ch == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if ch == rune(base[j]) {
				return false // Duplicate found
			}
		}
	}
	return true
}
