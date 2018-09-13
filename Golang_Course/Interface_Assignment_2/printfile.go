package main

import (
	"fmt"
	"io"
	"os"
)

// Type to implement the Writer interface in the standard library
type fileWriter struct{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file listed as an argument!")
		os.Exit(1)
	}

	file, _ := os.Open(os.Args[1])

	if file == nil {
		fmt.Println("Invalid file name!")
		os.Exit(1)
	}

	lw := fileWriter{}

	// First argument implements the Writer interface
	// Second argument implements the Reader interface
	io.Copy(lw, file)
}

// Custom function to satisfy Writer interface requirments
// and print out the file's contents to the console
func (fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
