package sort

import "testing"

func Test_swap(t *testing.T) {
	tests := []struct {
		name  string
		array *[]int
		idx1  int
		idx2  int
		want  []int
	}{
		{
			name:  "SwapTwo",
			array: &[]int{1, 2},
			idx1:  0,
			idx2:  1,
			want:  []int{2, 1},
		},
		{
			name:  "SwapItself",
			array: &[]int{1, 2},
			idx1:  0,
			idx2:  0,
			want:  []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.array, tt.idx1, tt.idx2)
			if !equalSlice(*tt.array, tt.want) {
				t.Errorf("swap() = %v, want %v", tt.array, tt.want)
			}
		})
	}
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
