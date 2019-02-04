package dfs

import (
	stack "../../../Data_Structures/Stack/Golang"
)

// TreeNode is the structure the binary tree is composed of
type TreeNode struct {
	value      string
	leftChild  *TreeNode
	rightChild *TreeNode
}

// NewBinaryTree returns the root node of a new binary tree with the structure:
//       A
//   B       E
// C   D   F   G
func NewBinaryTree() TreeNode {
	c := TreeNode{"C", nil, nil}
	d := TreeNode{"D", nil, nil}
	f := TreeNode{"F", nil, nil}
	g := TreeNode{"G", nil, nil}
	b := TreeNode{"B", &c, &d}
	e := TreeNode{"E", &f, &g}
	root := TreeNode{"A", &b, &e}
	return root
}

// PrintDFS prints out the values of the tree using the Depth-First Search approach
func PrintDFS(root TreeNode) string {
	result := ""

	visited := make(map[string]int)
	toVisit := stack.Stack{}
	toVisit.Push(&root)

	for {
		if toVisit.IsEmpty() {
			break
		}
		current := toVisit.Pop().(*TreeNode)

		_, present := visited[current.value]
		if present {
			// If the value is present in the map, there is a loop, and we skip it
		} else {
			visited[current.value] = 1
			result = result + current.value

			if current.rightChild != nil {
				toVisit.Push(current.rightChild)
			}
			if current.leftChild != nil {
				toVisit.Push(current.leftChild)
			}
		}
	}
	return result
}
