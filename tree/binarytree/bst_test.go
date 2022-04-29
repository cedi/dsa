package binarytree

import "testing"

func TestTreeNode_InsertBST(t *testing.T) {
	root := NewTreeNode(8)

	root.InsertBST(3)

	if root.left == nil || root.left.data != 3 {
		t.Errorf("TreeNode.Insert(3) expected left = 3, right = nil")
	}

	root.InsertBST(10)

	if root.right == nil || root.right.data != 10 {
		t.Errorf("TreeNode.Insert(10) expected left = 3, right = 10")
	}

	root.InsertBST(1)

	if root.left.left == nil || root.left.left.data != 1 {
		t.Errorf("TreeNode.Insert(1) expected left.left = 1")
	}

	root.InsertBST(6)

	if root.left.right == nil || root.left.right.data != 6 {
		t.Errorf("TreeNode.Insert(6) expected left.right = 6")
	}

	root.InsertBST(14)

	if root.right.right == nil || root.right.right.data != 14 {
		t.Errorf("TreeNode.Insert(14) expected right.right = 14")
	}

	root.InsertBST(9)

	if root.right.left == nil || root.right.left.data != 9 {
		t.Errorf("TreeNode.Insert(9) expected right.left = 9")
	}
}
