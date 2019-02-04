package dfs

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

	if root.leftChild.leftChild.value != "C" {
		t.Errorf("The root's left child must be 'C'. Got: %s", root.leftChild.leftChild.value)
	}

	if root.leftChild.rightChild.value != "D" {
		t.Errorf("Node must be 'D'. Got: %s", root.leftChild.rightChild.value)
	}

	if root.rightChild.value != "E" {
		t.Errorf("Node must be 'E'. Got: %s", root.rightChild.value)
	}

	if root.rightChild.leftChild.value != "F" {
		t.Errorf("Node must be 'F'. Got: %s", root.rightChild.leftChild.value)
	}

	if root.rightChild.rightChild.value != "G" {
		t.Errorf("Node must be 'G'. Got: %s", root.rightChild.rightChild.value)
	}
}

func TestPrintDFS(t *testing.T) {
	root := NewBinaryTree()
	s := PrintDFS(root)

	if s != "ABCDEFG" {
		t.Errorf("DFS is invalid. Expected 'ABCDEFG'. Got: %s", s)
	}

	s = PrintDFS(*root.leftChild)
	if s != "BCD" {
		t.Errorf("DFS is invalid. Expected 'BCD'. Got: %s", s)
	}

	s = PrintDFS(*root.rightChild)
	if s != "EFG" {
		t.Errorf("DFS is invalid. Expected 'EFG'. Got: %s", s)
	}
}
