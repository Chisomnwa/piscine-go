/*
Create the following function

func UltimateDivMod(a *int, b *int) {

}

UltimateDivmod should divide the deferenced value of a by the deferenced value of b.
- store the result of the division in the fmt which a points to
- Store the remainder of the division in the int which b points to
*/

package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a / *b
	mod := *a / *b
	*a = div
	*b = mod
}
