package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create and return a list of playing cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "JACK", "QUEEN", "KING"}

	// The underscore is to handle required unnused variable errors
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// Returning two values. Both of type deck
func deal(d deck, handSize int) (deck, deck) {
	// d[startIndexIncluding : upToNotIncluding]
	// If left side is blank, Go assumes the beginning
	// If right side is blank, Go assumes the end

	return d[:handSize], d[handSize:]
}
