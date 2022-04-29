package sort

// QuickSort
// best O(n log n) - needs one pass across the array and from there on it's divide and conquer
func QuickSort(array []int) []int {
	return quickSort(array, 0, len(array)-1, 0, "full")
}

func quickSort(array []int, low, high, pass int, orientation string) []int {
	pass++
	//fmt.Printf("Pass #%v - %v\n", pass, orientation)
	//fmt.Printf("Low: %v, High; %v\n", low, high)

	if low < high {
		pi := partition(array, low, high)
		//fmt.Printf("Pivot Idx: %v\n", pi)

		_ = quickSort(array, low, pi-1, pass, "left")
		//fmt.Printf("Left: %v \n", left)

		_ = quickSort(array, pi+1, high, pass, "right")
		//fmt.Printf("Right: %v \n", right)
	}

	return array
}

// place all elements smaller than pivot left and all greater than pivot to right
func partition(array []int, low, high int) int {

	// placing pivot at the right position in this array
	pivotIdx := high
	smallIdx := low

	//fmt.Printf("PivotIdx: %v Pivot: %v \n", pivotIdx, array[pivotIdx])

	for idx := low; idx < high; idx++ {
		//fmt.Printf("cmp %v < %v %v", array[idx], array[pivotIdx], array)
		if array[idx] < array[pivotIdx] {
			swap(&array, idx, smallIdx)
			//fmt.Printf(" -> swap(%v, %v) %v", idx, smallIdx, array)
			smallIdx++

		}

		//fmt.Printf("\n")

	}

	swap(&array, pivotIdx, smallIdx)
	//fmt.Printf("swap pivot into place -> swap(%v, %v) %v\n", pivotIdx, smallIdx, array)

	return smallIdx
}
