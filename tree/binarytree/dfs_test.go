package binarytree

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestTreeNode_InOrder(t *testing.T) {
	tests := []struct {
		name    string
		root    int
		inserts []int
		want    []int
	}{
		{
			name:    "insert one",
			root:    8,
			inserts: []int{3},
			want:    []int{3, 8},
		},
		{
			name:    "insert two",
			root:    8,
			inserts: []int{3, 10},
			want:    []int{3, 8, 10},
		},
		{
			name:    "insert multiple",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14, 9},
			want:    []int{1, 3, 6, 8, 9, 10, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTreeNode(tt.root)

			for _, item := range tt.inserts {
				tr.InsertBST(item)
			}

			got := tr.InOrder()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_PreOrder(t *testing.T) {
	tests := []struct {
		name    string
		root    int
		inserts []int
		want    []int
	}{
		{
			name:    "insert one",
			root:    8,
			inserts: []int{3},
			want:    []int{8, 3},
		},
		{
			name:    "insert two",
			root:    8,
			inserts: []int{3, 10},
			want:    []int{8, 3, 10},
		},
		{
			name:    "insert multiple",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14, 9},
			want:    []int{8, 3, 1, 6, 10, 9, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTreeNode(tt.root)

			for _, item := range tt.inserts {
				tr.InsertBST(item)
			}

			got := tr.PreOrder()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.PreOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_PostOrder(t *testing.T) {
	tests := []struct {
		name    string
		root    int
		inserts []int
		want    []int
	}{
		{
			name:    "insert one",
			root:    8,
			inserts: []int{3},
			want:    []int{3, 8},
		},
		{
			name:    "insert two",
			root:    8,
			inserts: []int{3, 10},
			want:    []int{3, 10, 8},
		},
		{
			name:    "insert multiple",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14, 9},
			want:    []int{1, 6, 3, 9, 14, 10, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTreeNode(tt.root)

			for _, item := range tt.inserts {
				tr.InsertBST(item)
			}

			got := tr.PostOrder()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.PreOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_GetListOfDepthsDFS(t *testing.T) {
	tests := []struct {
		name    string
		root    int
		inserts []int
		want    [][]int
	}{
		{
			name:    "One level",
			root:    8,
			inserts: []int{},
			want:    [][]int{{8}},
		},
		{
			name:    "Two level",
			root:    8,
			inserts: []int{3, 10},
			want:    [][]int{{8}, {3, 10}},
		},
		{
			name:    "Three level, imbalanced right not complete",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14},
			want:    [][]int{{8}, {3, 10}, {1, 6, 14}},
		},
		{
			name:    "Three level",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14, 9},
			want:    [][]int{{8}, {3, 10}, {1, 6, 14, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTreeNode(tt.root)

			for _, item := range tt.inserts {
				tr.InsertLevelOrder(item)
			}

			got := tr.GetListOfDepthsDFS()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.GetListOfDepthsDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTreeNode_GetListOfDepthsDFS1000(b *testing.B) {
	benchmarkTreeNode_GetListOfDepthsDFS(1000, b)
}

func BenchmarkTreeNode_GetListOfDepthsLevelOrderDFS1000(b *testing.B) {
	benchmarkTreeNode_GetListOfDepthsInsertLevelOrderDFS(1000, b)
}

func BenchmarkTreeNode_GetListOfDepthsDFS20000(b *testing.B) {
	benchmarkTreeNode_GetListOfDepthsDFS(20000, b)
}

func BenchmarkTreeNode_GetListOfDepthsLevelOrderDFS20000(b *testing.B) {
	benchmarkTreeNode_GetListOfDepthsInsertLevelOrderDFS(20000, b)
}

func benchmarkTreeNode_GetListOfDepthsDFS(i int, b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	tr := NewTreeNode(rand.Int())

	for idx := 0; idx < i; idx++ {
		tr.InsertBST(rand.Int())
	}

	for n := 0; n < b.N; n++ {
		tr.GetListOfDepthsDFS()
	}
}

func benchmarkTreeNode_GetListOfDepthsInsertLevelOrderDFS(i int, b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	tr := NewTreeNode(rand.Int())

	for idx := 0; idx < i; idx++ {
		tr.InsertLevelOrder(rand.Int())
	}

	for n := 0; n < b.N; n++ {
		tr.GetListOfDepthsDFS()
	}
}
