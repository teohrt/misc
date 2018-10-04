package main

import "fmt"

type Stack []int

func (s *Stack) push(v int) {
	*s = append(*s, v)
}

func (s *Stack) pop() int {
	if (s.isEmpty()) {
		fmt.Println(" The stack is already empty.")
		return -1;
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v;
}

func (s *Stack) peek() int {
	if (s.isEmpty()) {
		fmt.Println(" The stack is already empty.")
		return -1;
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0;
}