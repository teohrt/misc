package main

import "fmt"

func main() {
	stack := Stack{}

	str := "Reverse this string by using a stack!"
	fmt.Println(str)

	for i, _ := range str {
		stack.push(str[i])
	}

	for !stack.isEmpty() {
		fmt.Print(string(stack.pop()))
	}
	
	fmt.Println()


}