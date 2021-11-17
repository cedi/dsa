package sort

// MergeSort
// O(n log n) because I need log(n) thanks to divide and conquer, but n steps to merge
func MergeSort(array []int) []int {

	if len(array) < 2 {
		// no recurse
		return array
	}

	pivot := len(array) / 2

	left := MergeSort(array[:pivot])
	right := MergeSort(array[pivot:])

	//fmt.Printf("aLeft: %v, aRight: %v\n", left, right)

	sorted := []int{}

	leftIdx := 0
	leftLen := len(left)

	rightIdx := 0
	rightLen := len(right)

	for {
		if !(leftIdx < leftLen) {
			break
		}

		if !(rightIdx < rightLen) {
			break
		}

		itmLeft := left[leftIdx]
		itmRight := right[rightIdx]

		//fmt.Printf("iLeft: %v, iRight: %v\n", itmLeft, itmRight)

		if itmLeft < itmRight {
			sorted = append(sorted, itmLeft)
			leftIdx++
		} else {
			sorted = append(sorted, itmRight)
			rightIdx++
		}

		//fmt.Printf("Sorted: %v\n", sorted)
	}

	if leftIdx < leftLen {
		sorted = append(sorted, left[leftIdx:]...)
	}
	if rightIdx < rightLen {
		sorted = append(sorted, right[rightIdx:]...)
	}

	return sorted
}
