package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create and return a list of playing cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "JACK", "QUEEN", "KING"}

	// The underscores handle required unnused variable's errors
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// The function takes a reciever deck and prints all the cards
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

// Joining a slice of Strings!
// Converts a slice of deck to a slice of strings
// Converts the slice of string to a single string with a comma delimiter
// Returns that string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	// Arguemnts: (filename, byte slice of string, permsions number representing that anyone can read or write)
}

func newDeckFromFile(filename string) deck {
	// 'bs' is a byte slice from the input file
	bs, err := ioutil.ReadFile(filename)

	// If something goes wrong
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Convert byte slice to string.
	// Split on the comma delimiter and convert the string to a string slice
	// Convert that string slice to a deck and return it
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
