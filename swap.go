/*
Write a function that takes two pointers to an int (*int)
and swaps their content
*/

package piscine

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
