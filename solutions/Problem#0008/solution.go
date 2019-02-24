package main

import "fmt"

// Node of unival tree
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func main() {
	node := Node{1, &Node{1, nil, nil}, &Node{1, nil, nil}}
	fmt.Println(getUnivalCount(&node))
}

func getUnivalCount(root *Node) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftCount, isLeftUnival := getUnivalCount(root.Left)
	rightCount, isRightUnival := getUnivalCount(root.Right)
	if isLeftUnival && isRightUnival {
		if (root.Left != nil && root.Data != root.Left.Data) || (root.Right != nil && root.Data != root.Right.Data) {
			return leftCount + rightCount, false
		}
		return leftCount + rightCount + 1, true
	}
	return leftCount + rightCount, false
}
