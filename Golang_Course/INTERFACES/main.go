package main

import "fmt"

type bot interface {
	// Any type that has this function
	// is now also considered type bot!
	// Type bot can share functions
	// with other types this way
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Custom logic for generating an english greeting
	return "Hi there!"
}

// You don't need to have a reciever variable if it's not used!!
func (spanishBot) getGreeting() string {
	return "Hola!"
}
