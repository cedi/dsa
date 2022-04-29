package linkedlist

import (
	"testing"
)

func TestLinkedList_InsertBack(t *testing.T) {
	listFront := NewFromArgs(1, 2, 3)

	tests := []struct {
		name string
		list *LinkedList
		data int
		want *LinkedList
	}{
		{
			name: "InsertBack",
			list: listFront,
			data: 4,
			want: NewFromArgs(1, 2, 3, 4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.list

			l.InsertBack(tt.data)

			if l.String() != tt.want.String() {
				t.Errorf("[%s] LinkedList.InsertBack() = got %v, want %v", tt.name, l.String(), tt.want.String())
			}

			if l.Len() != tt.want.Len() {
				t.Errorf("[%s] LinkedList.InsertBack() = length %v, want %v", tt.name, l.Len(), tt.want.Len())
			}
		})
	}
}

func TestLinkedList_InsertFront(t *testing.T) {
	listFront := NewFromArgs(1, 2, 3)

	tests := []struct {
		name string
		list *LinkedList
		data int
		want *LinkedList
	}{
		{
			name: "InsertFront",
			list: listFront,
			data: 4,
			want: NewFromArgs(4, 1, 2, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.list

			l.InsertFront(tt.data)

			if l.String() != tt.want.String() {
				t.Errorf("[%s] LinkedList.InsertFront() = got %v, want %v", tt.name, l.String(), tt.want.String())
			}

			if l.Len() != tt.want.Len() {
				t.Errorf("[%s] LinkedList.InsertFront() = length %v, want %v", tt.name, l.Len(), tt.want.Len())
			}
		})
	}
}

func TestLinkedList_InsertAfter(t *testing.T) {
	listMiddle := NewFromArgs(1, 2, 3)
	listEnd := NewFromArgs(1, 2, 3)
	listFront := NewFromArgs(1, 2, 3)

	type args struct {
		node *Node
		data int
	}
	tests := []struct {
		name string
		list *LinkedList
		args args
		want *LinkedList
	}{
		{
			name: "InsertInMiddle",
			list: listMiddle,
			args: args{
				node: listMiddle.GetHead(),
				data: 4,
			},
			want: NewFromArgs(1, 4, 2, 3),
		},
		{
			name: "InsertBack",
			list: listEnd,
			args: args{
				node: listEnd.GetHead().Next.Next,
				data: 4,
			},
			want: NewFromArgs(1, 2, 3, 4),
		},
		{
			name: "InsertFront",
			list: listFront,
			args: args{
				node: nil,
				data: 4,
			},
			want: NewFromArgs(4, 1, 2, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.list

			l.InsertAfter(tt.args.node, tt.args.data)

			if l.String() != tt.want.String() {
				t.Errorf("[%s] LinkedList.InsertAfter() = got %v, want %v", tt.name, l.String(), tt.want.String())
			}

			if l.Len() != tt.want.Len() {
				t.Errorf("[%s] LinkedList.InsertAfter() = length %v, want %v", tt.name, l.Len(), tt.want.Len())
			}
		})
	}
}

func TestLinkedList_InsertBefore(t *testing.T) {
	listMiddle := NewFromArgs(1, 2, 3)
	listFront := NewFromArgs(1, 2, 3)
	listNil := NewFromArgs(1, 2, 3)

	type args struct {
		node *Node
		data int
	}
	tests := []struct {
		name string
		list *LinkedList
		args args
		want *LinkedList
	}{
		{
			name: "InsertInMiddle",
			list: listMiddle,
			args: args{
				node: listMiddle.GetHead().Next,
				data: 4,
			},
			want: NewFromArgs(1, 4, 2, 3),
		},
		{
			name: "InsertFront",
			list: listFront,
			args: args{
				node: listMiddle.GetHead(),
				data: 4,
			},
			want: NewFromArgs(4, 1, 2, 3),
		},
		{
			name: "InsertNil",
			list: listNil,
			args: args{
				node: nil,
				data: 4,
			},
			want: NewFromArgs(4, 1, 2, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.list

			l.InsertBefore(tt.args.node, tt.args.data)

			if l.String() != tt.want.String() {
				t.Errorf("[%s] LinkedList.InsertBefore() = got %v, want %v", tt.name, l.String(), tt.want.String())
			}

			if l.Len() != tt.want.Len() {
				t.Errorf("[%s] LinkedList.InsertBefore() = length %v, want %v", tt.name, l.Len(), tt.want.Len())
			}
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	listMiddle := NewFromArgs(1, 2, 3)
	listEnd := NewFromArgs(1, 2, 3)
	listFront := NewFromArgs(1, 2, 3)

	tests := []struct {
		name string
		list *LinkedList
		node *Node
		want *LinkedList
	}{
		{
			name: "DeleteMiddle",
			list: listMiddle,
			node: listMiddle.GetHead().Next,
			want: NewFromArgs(1, 3),
		},
		{
			name: "DeleteLast",
			list: listEnd,
			node: listEnd.GetHead().Next.Next,
			want: NewFromArgs(1, 2),
		},
		{
			name: "DeleteFirst",
			list: listFront,
			node: listFront.GetHead(),
			want: NewFromArgs(2, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.list

			l.Delete(tt.node)

			if l.String() != tt.want.String() {
				t.Errorf("[%s] LinkedList.InsertAfter() = got %v, want %v", tt.name, l.String(), tt.want.String())
			}

			if l.Len() != tt.want.Len() {
				t.Errorf("[%s] LinkedList.InsertAfter() = length %v, want %v", tt.name, l.Len(), tt.want.Len())
			}
		})
	}
}

func BenchmarkInsertFront100(b *testing.B) {
	benchmarkInsertFront(100, b)
}

func BenchmarkInsertFront1000(b *testing.B) {
	benchmarkInsertFront(1000, b)
}

func BenchmarkInsertFront10000(b *testing.B) {
	benchmarkInsertFront(10000, b)
}

func BenchmarkInsertFront100000(b *testing.B) {
	benchmarkInsertFront(100000, b)
}

func BenchmarkInsertBack100(b *testing.B) {
	benchmarkInsertBack(100, b)
}

func BenchmarkInsertBack1000(b *testing.B) {
	benchmarkInsertBack(1000, b)
}

func BenchmarkInsertBack10000(b *testing.B) {
	benchmarkInsertBack(10000, b)
}

func BenchmarkInsertBack100000(b *testing.B) {
	benchmarkInsertBack(100000, b)
}

func benchmarkInsertFront(elements int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		list := New()
		for i := 1; i <= elements; i++ {
			list.InsertFront(i)
		}
	}
}

func benchmarkInsertBack(elements int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		list := New()
		for i := 1; i <= elements; i++ {
			list.InsertBack(i)
		}
	}
}
