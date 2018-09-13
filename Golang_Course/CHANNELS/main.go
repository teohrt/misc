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

	// Each requests must wait for the previous one to complete
	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
