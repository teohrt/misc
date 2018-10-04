package main

import "fmt"

func main() {
	stack := Stack{}

	fmt.Println(stack.isEmpty())
	stack.push(10)
	stack.push(20)
	fmt.Println(stack.isEmpty())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.isEmpty())
}