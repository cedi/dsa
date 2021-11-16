package linkedlist

// DetectLoop tries to detect a loop and returns true if it found one and false if not
func DetectLoop(list *LinkedList) bool {
	iterSlow := list.GetHead()
	iterFast := list.GetHead()

	for iterSlow != nil && iterFast != nil && iterFast.Next != nil {
		iterSlow = iterSlow.Next
		iterFast = iterFast.Next.Next

		if iterSlow == iterFast {
			return true
		}
	}

	return false
}
