package main

import "fmt"

func main() {
	// This is going to be a slice of type string
	cards := []string{"Ace of Diamonds", newCard()}

	// This does not modify the slice. It returns a new one!
	cards = append(cards, "Six of Spades")

	// Iterate over cards
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// Seperate function returning a string
func newCard() string {
	return "Five of Diamonds"
}
