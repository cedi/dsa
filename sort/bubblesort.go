package sort

// BubbleSort
// best O(n) worst O(n^2) depending on the amount of rounds needed to sort everything
func BubbleSort(array []int) []int {
	for cntr := 1; ; cntr++ {
		//fmt.Printf("Pass #%v\n", cntr)
		swapped := false
		for idx := 0; idx < len(array)-1; idx++ {
			//fmt.Printf("cmp %v > %v %v", array[idx], array[idx+1], array)
			if array[idx] > array[idx+1] {
				swap(&array, idx, idx+1)
				//fmt.Printf(" %v -> swap!", array)
				swapped = true
			}
			//fmt.Printf("\n")
		}

		if !swapped {
			break
		}
	}

	return array
}
