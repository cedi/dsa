package linkedlist

import (
	"math/rand"
	"testing"
)

func TestLinkedList_InsertionSort(t *testing.T) {
	tests := []struct {
		name string
		list *LinkedList
		want *LinkedList
	}{
		{
			name: "OneOff",
			list: NewFromArgs(2, 1, 3),
			want: NewFromArgs(1, 2, 3),
		},
		{
			name: "Reversed",
			list: NewFromArgs(3, 2, 1),
			want: NewFromArgs(1, 2, 3),
		},
		{
			name: "TwoEqual",
			list: NewFromArgs(3, 2, 2, 1),
			want: NewFromArgs(1, 2, 2, 3),
		},
		{
			name: "Longer",
			list: NewFromArgs(3, 2, 1, 4, 5, 6),
			want: NewFromArgs(1, 2, 3, 4, 5, 6),
		},
		{
			name: "Randomer",
			list: NewFromArgs(3, 1, 5, 2, 10, 6),
			want: NewFromArgs(1, 2, 3, 5, 6, 10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertionSort(tt.list)

			if got.String() != tt.want.String() {
				t.Errorf("[%s] LinkedList.InsertFront() = got %v, want %v", tt.name, got.String(), tt.want.String())
			}

			if got.Len() != tt.want.Len() {
				t.Errorf("[%s] LinkedList.InsertFront() = length %v, want %v", tt.name, got.Len(), tt.want.Len())
			}
		})
	}
}

func BenchmarkSort10(b *testing.B) {
	benchmarkSort(10, b)
}
func BenchmarkSort100(b *testing.B) {
	benchmarkSort(100, b)
}
func BenchmarkSort1000(b *testing.B) {
	benchmarkSort(1000, b)
}
func BenchmarkSort10000(b *testing.B) {
	benchmarkSort(10000, b)
}

func benchmarkSort(elements int, b *testing.B) {
	list := New()
	for i := 1; i <= elements; i++ {
		list.InsertFront(rand.Intn(elements))
	}

	for n := 0; n < b.N; n++ {
		InsertionSort(list)
	}
}
