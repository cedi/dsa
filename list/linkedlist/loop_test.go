package linkedlist

import (
	"math/rand"
	"testing"
)

func TestDetectLoop(t *testing.T) {
	bigLoop := NewFromArgs(1, 2, 3, 4)
	bigLoop.GetLast().Next = bigLoop.GetHead()

	smallLoop := NewFromArgs(1, 2, 3, 4)
	smallLoop.GetLast().Next = smallLoop.GetHead().Next.Next

	tests := map[string]struct {
		list *LinkedList
		want bool
	}{
		"NoLoop": {
			list: NewFromArgs(1, 2, 3, 4),
			want: false,
		},
		"Bigloop": {
			list: bigLoop,
			want: true,
		},
		"SmalLoop": {
			list: smallLoop,
			want: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := DetectLoop(tt.list); got != tt.want {
				t.Errorf("DetectLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDetectLoop100(b *testing.B) {
	benchmarkDetectLoop(100, b)
}
func BenchmarkDetectLoop1000(b *testing.B) {
	benchmarkDetectLoop(1000, b)
}
func BenchmarkDetectLoop10000(b *testing.B) {
	benchmarkDetectLoop(10000, b)
}
func BenchmarkDetectLoop100000(b *testing.B) {
	benchmarkDetectLoop(100000, b)
}

func benchmarkDetectLoop(elements int, b *testing.B) {
	list := New()
	for i := 1; i <= elements; i++ {
		list.InsertFront(rand.Intn(elements))
	}

	list.GetLast().Next = list.GetHead()

	for n := 0; n < b.N; n++ {
		DetectLoop(list)
	}
}
