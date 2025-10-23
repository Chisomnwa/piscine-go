package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var burgerTime food
	var chipsTime food
	var nuggetsTime food

	burgerTime.preptime = 15
	chipsTime.preptime = 10
	nuggetsTime.preptime = 12

	// Compare the input order with each menu
	if order == "burger" {
		return burgerTime.preptime
	} else if order == "chips" {
		return chipsTime.preptime
	} else if order == "nuggets" {
		return nuggetsTime.preptime
	}

	return 404
}
