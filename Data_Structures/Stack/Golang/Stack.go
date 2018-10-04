package main

import "fmt"

type Stack []byte

func (s *Stack) push(v byte) {
	*s = append(*s, v)
}

func (s *Stack) pop() byte {
	if (s.isEmpty()) {
		fmt.Println(" The stack is already empty.")
		return ' ';
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v;
}

func (s *Stack) peek() byte {
	if (s.isEmpty()) {
		fmt.Println(" The stack is already empty.")
		return ' ';
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0;
}