package linkedlist

import (
	"fmt"
	"testing"
)

var l LinkedList

func TestIsEmpty(t *testing.T) {
	if !l.IsEmpty() {
		t.Errorf("This list should be empty")
	}
}

func TestAppendSize(t *testing.T) {
	l.Append(0)
	if l.size != 1 {
		t.Errorf("This list should have one element")
	}
	l.Append(0)
	if l.size != 2 {
		t.Errorf("This list should have two elements")
	}
}

func TestRemoveHeadValue(t *testing.T) {
	values := []Item{[]int{1, 2, 3}, 0, map[string]int{"testing": 123}}

	clearOutLinkedList()

	for _, value := range values {
		l.Append(value)
	}

	var r = l.RemoveHead()
	if r.([]int)[0] != 1 && r.([]int)[1] != 2 && r.([]int)[2] != 3 {
		fmt.Println(r.([]int)[0])
		t.Errorf("Slice of integers wasn't returned correctly. Got: %d", r.([]int)[0])
	}

	r = l.RemoveHead()
	if r.(int) != 0 {
		t.Errorf("Int wasn't returned correctly. Got: %s", r)
	}

	r = l.RemoveHead()
	if r.(map[string]int)["testing"] != 123 {
		t.Errorf("Map wasn't returned correctly. Got: %s", r)
	}

}

func clearOutLinkedList() {
	if !l.IsEmpty() {
		for {
			if l.size == 0 {
				break
			}
			l.RemoveHead()
		}
	}
}
