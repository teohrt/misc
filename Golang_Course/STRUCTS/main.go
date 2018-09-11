package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// You don't need to assign a name to a struct here
	contactInfo
}

func main() {
	// There are a few different ways to initialize
	// trace := person{"Trace", "Ohrt"}

	// var trace person
	// trace.firstName = "Trace"
	// trace.lastName = "Ohrt"

	trace := person{
		firstName: "Trace",
		lastName:  "Ohrt",
		contactInfo: contactInfo{
			email:   "teohrt18@gmail.com",
			zipCode: 50010,
		},
	}

	// YOU DON'T NEED TO DO THIS IN GO
	// tracePointer := &trace
	trace.updateName("TRACEEEEE")
	trace.print()

}

// Pointers!
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
