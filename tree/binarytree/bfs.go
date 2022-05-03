package binarytree

func (t *TreeNode) LevelOrderTraversal(fn TraversalFunc) {
	treeQueue := make([]*TreeNode, 0)

	treeQueue = append(treeQueue, t)
	curLevel := 0

Exit:
	for len(treeQueue) > 0 {
		for levelSize := len(treeQueue); levelSize != 0; levelSize-- {
			curNode := treeQueue[0]
			treeQueue = treeQueue[1:]

			ok := fn(curLevel, curNode)
			if !ok {
				break Exit
			}

			if curNode.left != nil {
				treeQueue = append(treeQueue, curNode.left)
			}

			if curNode.right != nil {
				treeQueue = append(treeQueue, curNode.right)
			}
		}
		curLevel++
	}
}

func (t *TreeNode) InsertLevelOrder(data int) *TreeNode {
	newNode := NewTreeNode(data)

	t.LevelOrderTraversal(func(_ int, node *TreeNode) bool {
		if node.left == nil {
			node.left = newNode
			return false
		}

		if node.right == nil {
			node.right = newNode
			return false
		}

		return true
	})

	return newNode
}

func (t *TreeNode) LevelOrder() []int {
	res := make([]int, 0)

	t.LevelOrderTraversal(func(_ int, node *TreeNode) bool {
		res = append(res, node.data)
		return true
	})

	return res
}

func (t *TreeNode) GetListOfDepths() [][]int {
	depth := t.Depth()

	listOfDepthLevels := make([][]int, depth)

	t.LevelOrderTraversal(func(level int, node *TreeNode) bool {
		levelList := listOfDepthLevels[level]
		levelList = append(levelList, node.data)
		listOfDepthLevels[level] = levelList
		return true
	})

	return listOfDepthLevels
}
