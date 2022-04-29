package binarytree

type LevelOrderTraversalFunc func(*TreeNode, int) bool

func (t *TreeNode) LevelOrderTraversal(fn LevelOrderTraversalFunc) {
	treeQueue := make([]*TreeNode, 0)

	treeQueue = append(treeQueue, t)
	curLevel := 0

Exit:
	for len(treeQueue) > 0 {
		for levelSize := len(treeQueue); levelSize != 0; levelSize-- {
			curNode := treeQueue[0]
			treeQueue = treeQueue[1:]

			ok := fn(curNode, curLevel)
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

	t.LevelOrderTraversal(func(node *TreeNode, _ int) bool {
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

	t.LevelOrderTraversal(func(node *TreeNode, _ int) bool {
		res = append(res, node.data)
		return true
	})

	return res
}

func (t *TreeNode) GetListOfDepths() [][]int {
	depth := t.Depth()

	listOfDepthLevels := make([][]int, depth)

	t.LevelOrderTraversal(func(node *TreeNode, level int) bool {
		levelList := listOfDepthLevels[level]
		levelList = append(levelList, node.data)
		listOfDepthLevels[level] = levelList
		return true
	})

	return listOfDepthLevels
}
