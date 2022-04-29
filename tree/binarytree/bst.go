package binarytree

func (t *TreeNode) InsertBST(data int) *TreeNode {
	if data <= t.data {
		if t.left == nil {
			t.left = NewTreeNode(data)
			return t.left
		}

		return t.left.InsertBST(data)
	}

	if t.right == nil {
		t.right = NewTreeNode(data)
		return t.left
	}

	return t.right.InsertBST(data)
}
