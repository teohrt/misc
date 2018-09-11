package main

import (
	"fmt"
)

func main() {
	// Much like structs there are a few different ways to create maps

	// var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
	}
	colors["white"] = "#ffffff"
	delete(colors, "red")

	fmt.Println(colors)

}
