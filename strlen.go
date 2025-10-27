/*
Write a function that counts the runes of a string and that returns that count.
*/
package piscine

func StrLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}
