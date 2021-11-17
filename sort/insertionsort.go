package sort

// InsertionSort
// always O(n^2) because of nested loop
func InsertionSort(array []int) []int {
	for idx := 1; idx < len(array); idx++ {
		for cmp := idx; cmp > 0; cmp-- {
			//fmt.Printf("cmp %v < %v", array[cmp], array[cmp-1])
			if array[cmp] < array[cmp-1] {
				//fmt.Printf(" -> swap!")
				swap(&array, cmp, cmp-1)
			}
			//fmt.Printf("\n")
		}
	}
	return array
}
