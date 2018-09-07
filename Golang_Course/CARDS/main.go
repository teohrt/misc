package main

func main() {
	// This is going to be a slice of type string
	cards := deck{"Ace of Diamonds", newCard()}

	// This does not modify the slice. It returns a new one!
	cards = append(cards, "Six of Spades")

	// Print cards
	cards.print()
}

// Seperate function returning a string
func newCard() string {
	return "Five of Diamonds"
}
