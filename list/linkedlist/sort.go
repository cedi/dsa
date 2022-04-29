package linkedlist

func InsertionSort(list *LinkedList) *LinkedList {
	sortedList := New()

	for iter := list.GetHead(); iter != nil; iter = iter.Next {
		insertSorted(sortedList, iter.Data)
	}

	return sortedList
}

func insertSorted(sortedList *LinkedList, data int) {
	if sortedList.head == nil {
		sortedList.InsertFront(data)
		return
	}

	for iter := sortedList.GetHead(); iter != nil; iter = iter.Next {
		if data <= iter.Data {
			sortedList.InsertBefore(iter, data)
			break
		}

		if iter.Next == nil {
			sortedList.InsertAfter(iter, data)
			break
		}

		if data < iter.Next.Data {
			sortedList.InsertAfter(iter, data)
			break
		}
	}
}
