package binarytree

import (
	"reflect"
	"testing"
)

func TestTreeNode_Depth(t *testing.T) {
	tests := []struct {
		name    string
		root    int
		inserts []int
		want    int
	}{
		{
			name:    "One level",
			root:    8,
			inserts: []int{},
			want:    1,
		},
		{
			name:    "Two level",
			root:    8,
			inserts: []int{3, 10},
			want:    2,
		},
		{
			name:    "Three levels, balanced",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14, 9},
			want:    3,
		},
		{
			name:    "Three levels, imbalanced left heavy",
			root:    8,
			inserts: []int{3, 10, 1, 6},
			want:    3,
		},
		{
			name:    "Three levels, imbalanced right heavy",
			root:    8,
			inserts: []int{3, 10, 14, 9},
			want:    3,
		},
		{
			name:    "Four levels, imbalanced left heavy",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14, 9, 5},
			want:    4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTreeNode(tt.root)

			for _, item := range tt.inserts {
				tr.InsertBST(item)
			}

			got := tr.Depth()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.Depth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_Print(t *testing.T) {
	tests := []struct {
		name    string
		root    int
		inserts []int
		want    string
	}{
		{
			name:    "insert one",
			root:    8,
			inserts: []int{3},
			want:    "[3 8]",
		},
		{
			name:    "insert two",
			root:    8,
			inserts: []int{3, 10},
			want:    "[3 8 10]",
		},
		{
			name:    "insert multiple",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14, 9},
			want:    "[1 3 6 8 9 10 14]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTreeNode(tt.root)

			for _, item := range tt.inserts {
				tr.InsertBST(item)
			}

			got := tr.String()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
