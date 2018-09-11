package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// trace := person{"Trace", "Ohrt"}
	// trace := person{firstName: "Trace", lastName: "Ohrt"}
	var trace person
	trace.firstName = "Trace"
	trace.lastName = "Ohrt"
	fmt.Println(trace)
	fmt.Printf("%+v\n", trace)

}
