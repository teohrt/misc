package main

import "fmt"

func main() {
	// Much like structs there are a few different ways to create maps

	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
	}
	// Adding a key value pair:
	colors["white"] = "#ffffff"

	// Deleting a pair:
	// delete(colors, "red")

	printMap(colors)

}

// Iterating over a map
func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
