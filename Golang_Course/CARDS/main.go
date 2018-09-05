package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := newCard()

	fmt.Println(card)
}

// Seperate function returning a string
func newCard() string {
	return "Five of Diamonds"
}
