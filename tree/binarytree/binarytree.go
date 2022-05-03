package binarytree

import "fmt"

type TraversalFunc func(int, *TreeNode) bool

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}

func NewTreeNode(data int) *TreeNode {
	return &TreeNode{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (t *TreeNode) Depth() int {
	left := 0
	right := 0

	if t.left == nil && t.right == nil {
		return 1
	}

	if t.left != nil {
		left = 1 + t.left.Depth()
	}

	if t.right != nil {
		right = 1 + t.right.Depth()
	}

	if left >= right {
		return left
	}

	return right
}

func (t *TreeNode) String() string {
	inOrder := t.InOrder()

	return fmt.Sprintf("%v", inOrder)
}
