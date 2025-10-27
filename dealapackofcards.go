package piscine

import "fmt"

func DealAPackofCards(deck []int) {
	cardsPerPlayer := len(deck) / 4 // 12 / 4 = 3

	for player := 0; player < 4; player++ {
		fmt.Printf("Player %d: ", player+1)

		// Print each player's 3 cards
		for i := 0; i < cardsPerPlayer; i++ {
			cardIndex := player*cardsPerPlayer + i
			fmt.Printf("%d", deck[cardIndex])

			// Add comma and space after each card except the last
			if i < cardsPerPlayer-1 {
				fmt.Print(", ")
			}
		}
		fmt.Printf("/n")
	}
}
