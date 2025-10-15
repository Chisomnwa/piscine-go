
package piscine

func IsUpper(s string) bool {
	for _, letter := range s {              // loop through each rune (character)
		if letter < 'A' || letter > 'Z' { 	// if any character is not uppercase
			return false
		}
	}
	return true 							// if all characters were uppercase
}
