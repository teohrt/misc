package main

import (
	"fmt"
	"net/http"
)

// Website checking program
func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://github.com",
		"http://facebook.com",
	}

	// Create a channel to pass into our function
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// THE WHOLE LOOP IS NOT EXCECUTED ALL AT ONCE
	// <-c is a blocking call
	// This means that when the program gets to it, it temporarily pauses
	// And waits for another go routine to send a message over the channel
	// This is so cool!
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- "DOWN! : " + link
		return
	}
	c <- "UP AND RUNNING!: " + link
}
