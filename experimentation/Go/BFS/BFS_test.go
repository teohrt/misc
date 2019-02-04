package bfs

import (
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	root := NewBinaryTree()
	if root.value != "A" {
		t.Errorf("Binary Tree Root node should be 'A'. Got: %s", root.value)
	}

	if root.leftChild.value != "B" {
		t.Errorf("The root's left child must be 'B'. Got: %s", root.leftChild.value)
	}

	if root.rightChild.value != "C" {
		t.Errorf("The root's left child must be 'C'. Got: %s", root.rightChild.value)
	}

	if root.leftChild.leftChild.value != "D" {
		t.Errorf("Node must be 'D'. Got: %s", root.leftChild.leftChild.value)
	}

	if root.leftChild.rightChild.value != "E" {
		t.Errorf("Node must be 'E'. Got: %s", root.leftChild.rightChild.value)
	}

	if root.rightChild.leftChild.value != "F" {
		t.Errorf("Node must be 'F'. Got: %s", root.rightChild.leftChild.value)
	}

	if root.rightChild.rightChild.value != "G" {
		t.Errorf("Node must be 'G'. Got: %s", root.rightChild.rightChild.value)
	}
}

func TestPrintBFS(t *testing.T) {
	root := NewBinaryTree()
	s := PrintBFS(root)

	if s != "ABCDEFG" {
		t.Errorf("BFS is invalid. Expected 'ABCDEFG'. Got: %s", s)
	}

	s = PrintBFS(*root.leftChild)
	if s != "BDE" {
		t.Errorf("BFS is invalid. Expected 'BDE'. Got: %s", s)
	}

	s = PrintBFS(*root.rightChild)
	if s != "CFG" {
		t.Errorf("BFS is invalid. Expected 'CFG'. Got: %s", s)
	}

}
