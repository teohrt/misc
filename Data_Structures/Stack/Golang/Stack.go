package stack

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

// Item is a "generic type" our Stack will store
type Item generic.Type

// Stack is fundamentally just a slice of items with extra functionality.
type Stack []Item

// Push appends an item to the stack
func (s *Stack) Push(v Item) {
	*s = append(*s, v)
}

// Pop removes and returns the item at the end of the stack
func (s *Stack) Pop() Item {
	if s.IsEmpty() {
		fmt.Println(" The stack is already empty.")
		return ' '
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

// Peek returns the Item at the end of the stack
func (s *Stack) Peek() Item {
	if s.IsEmpty() {
		fmt.Println(" The stack is already empty.")
		return ' '
	}
	return (*s)[len(*s)-1]
}

// IsEmpty returns a bool representing whether or not the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
