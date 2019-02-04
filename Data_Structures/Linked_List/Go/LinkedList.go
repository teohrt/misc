package linkedlist

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item is a "generic type" our Node store of value
type Item generic.Type

// Node is the base structure the LinkedList is comprised of
// Stores the content value, and a pointer to the next node in the list
type Node struct {
	content Item
	next    *Node
}

// The LinkedList holds a pointer to its head node, an int storing its size, and a Mutex object to make it concurrency safe
type LinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

// IsEmpty returns a bool representing whether the list is empty or not
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// Append adds a node to the end of the LinkedList
func (l *LinkedList) Append(item Item) {
	l.lock.Lock()
	node := Node{item, nil}
	// If empty
	if l.head == nil {
		l.head = &node
	} else {
		// Iterate to the end
		last := l.head
		for {
			if last.next == nil {
				break
			} else {
				last = last.next
			}
		}
		last.next = &node
	}
	l.size++
	l.lock.Unlock()
}

// RemoveHead deletes the first node in the LinkedList and returns its value
func (l *LinkedList) RemoveHead() Item {
	l.lock.Lock()
	v := l.head.content
	if l.size <= 1 {
		l.head = nil
		l.size = 0
	} else {
		l.head = l.head.next
		l.size--
	}
	l.lock.Unlock()
	return v
}
