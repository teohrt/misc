package main

func main() {
	// This is going to be a slice of type string
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// Multiple return values!
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// Joining a slice of strings!
	// fmt.Println(cards.toString())

	// Saving data to the hard drive
	//cards.saveToFile("my_cards")

	// Reading from the hard drive
	// cards := newDeckFromFile("my_cards")
	// cards.print()
}
