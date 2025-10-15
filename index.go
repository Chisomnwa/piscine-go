package piscine

func Index(s string, toFind string) int {
	// Loop through the main string `s`
	for i := 0; i <= len(s)-len(toFind); i++ {
		// Check if the substring at i matches `toFind`
		if s[i:i+len(toFind)] == toFind {
			return i // Return the starting index of the match
		}
	}
	return -1 // Return -1 if not allowed
}
