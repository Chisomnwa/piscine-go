package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	charCount := 0

	for i := 0; i < len(str); i++ {
		char := str[i]

		if char == ' ' {
			continue
		}

		result += string(char)
		charCount++

		if charCount == 5 {
			charCount = 0
			i++

			hasNextGoodChar := false
			for j := i + 1; j < len(str); j++ {
				if str[j] != ' ' {
					hasNextGoodChar = true
					break
				}
			}

			if hasNextGoodChar {
				result += " "
			}
		}
	}

	result += "\n"
	return result
}
