package sort

func swap(array *[]int, idx1, idx2 int) {
	intermediate := (*array)[idx1]
	(*array)[idx1] = (*array)[idx2]
	(*array)[idx2] = intermediate
}
