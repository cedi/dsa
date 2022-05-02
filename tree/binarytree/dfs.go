package binarytree

func (t *TreeNode) InOrder() []int {
	res := make([]int, 0)

	if t.left != nil {
		left := t.left.InOrder()
		res = append(res, left...)
	}

	res = append(res, t.data)

	if t.right != nil {
		right := t.right.InOrder()
		res = append(res, right...)
	}

	return res
}

func (t *TreeNode) PreOrder() []int {
	res := make([]int, 0)

	res = append(res, t.data)

	if t.left != nil {
		left := t.left.PreOrder()
		res = append(res, left...)
	}

	if t.right != nil {
		right := t.right.PreOrder()
		res = append(res, right...)
	}

	return res
}

func (t *TreeNode) PostOrder() []int {
	res := make([]int, 0)

	if t.left != nil {
		left := t.left.PostOrder()
		res = append(res, left...)
	}

	if t.right != nil {
		right := t.right.PostOrder()
		res = append(res, right...)
	}

	res = append(res, t.data)

	return res
}

func (t *TreeNode) GetListOfDepthsDFS() [][]int {
	depth := t.Depth()

	listOfDepthLevels := make([][]int, depth)

	listOfDepthLevels = t.getListOfDepthsDFS(0, listOfDepthLevels)

	return listOfDepthLevels
}

func (t *TreeNode) getListOfDepthsDFS(currentDepth int, listOfDepth [][]int) [][]int {

	curListOfDepth := listOfDepth[currentDepth]
	curListOfDepth = append(curListOfDepth, t.data) // pre-order traverse
	listOfDepth[currentDepth] = curListOfDepth

	if t.left != nil {
		t.left.getListOfDepthsDFS(1+currentDepth, listOfDepth)
	}

	if t.right != nil {
		t.right.getListOfDepthsDFS(1+currentDepth, listOfDepth)
	}

	return listOfDepth
}
