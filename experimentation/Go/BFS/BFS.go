package bfs

import (
	ll "../../../Data_Structures/Linked_List/Go"
)

// TreeNode is the structure the binary tree is composed of
type TreeNode struct {
	value      ll.Item
	leftChild  *TreeNode
	rightChild *TreeNode
}

/*
	NewBinaryTree returns the root node (in this case, A) of a new binary tree
	Tree Structure:
	        A
	     B     C
       D   E F   G
*/
func NewBinaryTree() TreeNode {
	d := TreeNode{"D", nil, nil}
	e := TreeNode{"E", nil, nil}
	f := TreeNode{"F", nil, nil}
	g := TreeNode{"G", nil, nil}
	b := TreeNode{"B", &d, &e}
	c := TreeNode{"C", &f, &g}
	root := TreeNode{"A", &b, &c}
	return root
}

// PrintBFS prints out the values of the table using the bredth first search approach
func PrintBFS(root TreeNode) string {
	result := ""

	visited := make(map[string]int)
	var toVisit ll.LinkedList
	toVisit.Append(&root)

	for {
		if toVisit.IsEmpty() {
			break
		}
		current := toVisit.RemoveHead().(*TreeNode)

		_, present := visited[current.value.(string)]
		if present {
			// If the value is present in the map, there is a loop, and we skip it
		} else {
			visited[current.value.(string)] = 1
			result = result + current.value.(string)

			if current.leftChild != nil {
				toVisit.Append(current.leftChild)
			}
			if current.rightChild != nil {
				toVisit.Append(current.rightChild)
			}
		}
	}
	return result
}
