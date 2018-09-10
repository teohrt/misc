package main

func main() {
	// This is going to be a slice of type string
	cards := newDeck()

	// Multiple return values!
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

}
