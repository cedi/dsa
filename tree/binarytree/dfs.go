package binarytree

func (t *TreeNode) InOrder() []int {
	res := make([]int, 0)

	t.InOrderTraversal(func(depth int, node *TreeNode) bool {
		res = append(res, node.data)
		return true
	})

	return res
}

func (t *TreeNode) PreOrder() []int {
	res := make([]int, 0)

	t.PreOrderTraversal(func(depth int, node *TreeNode) bool {
		res = append(res, node.data)
		return true
	})

	return res
}

func (t *TreeNode) PostOrder() []int {
	res := make([]int, 0)

	t.PostOrderTraversal(func(depth int, node *TreeNode) bool {
		res = append(res, node.data)
		return true
	})

	return res
}

func (t *TreeNode) InOrderTraversal(traversalFunc TraversalFunc) {
	t.inOrderTraversal(0, traversalFunc)
}

func (t *TreeNode) inOrderTraversal(depth int, traversalFunc TraversalFunc) {
	if t.left != nil {
		t.left.inOrderTraversal(depth+1, traversalFunc)
	}

	traversalFunc(depth, t)

	if t.right != nil {
		t.right.inOrderTraversal(depth+1, traversalFunc)
	}
}

func (t *TreeNode) PreOrderTraversal(traversalFunc TraversalFunc) {
	t.preOrderTraversal(0, traversalFunc)
}

func (t *TreeNode) preOrderTraversal(depth int, traversalFunc TraversalFunc) {
	traversalFunc(depth, t)

	if t.left != nil {
		t.left.preOrderTraversal(depth+1, traversalFunc)
	}

	if t.right != nil {
		t.right.preOrderTraversal(depth+1, traversalFunc)
	}
}

func (t *TreeNode) PostOrderTraversal(traversalFunc TraversalFunc) {
	t.postOrderTraversal(0, traversalFunc)
}

func (t *TreeNode) postOrderTraversal(depth int, traversalFunc TraversalFunc) {
	if t.left != nil {
		t.left.postOrderTraversal(depth+1, traversalFunc)
	}

	if t.right != nil {
		t.right.postOrderTraversal(depth+1, traversalFunc)
	}

	traversalFunc(depth, t)
}

func (t *TreeNode) GetListOfDepthsDFS() [][]int {
	depth := t.Depth()

	listOfDepthLevels := make([][]int, depth)

	t.PreOrderTraversal(func(currentDepth int, node *TreeNode) bool {
		curListOfDepth := listOfDepthLevels[currentDepth]
		curListOfDepth = append(curListOfDepth, node.data)
		listOfDepthLevels[currentDepth] = curListOfDepth
		return true
	})

	return listOfDepthLevels
}
