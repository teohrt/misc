package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// trace := person{"Trace", "Ohrt"}

	// var trace person
	// trace.firstName = "Trace"
	// trace.lastName = "Ohrt"

	trace := person{
		firstName: "Trace",
		lastName:  "Ohrt",
		contact: contactInfo{
			email:   "teohrt18@gmail.com",
			zipCode: 50010,
		},
	}

	fmt.Println(trace)
	fmt.Printf("%+v\n", trace)

}
