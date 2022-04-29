package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "Short",
			array: []int{3, 1, 2, 5, 2, 5, 6},
			want:  []int{1, 2, 2, 3, 5, 5, 6},
		},
		{
			name:  "One",
			array: []int{3},
			want:  []int{3},
		},
		{
			name:  "None",
			array: []int{},
			want:  []int{},
		},
		{
			name:  "InOrder",
			array: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Reversed",
			array: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "OneOff",
			array: []int{1, 3, 2, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
