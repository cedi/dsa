package binarytree

import (
	"reflect"
	"testing"
)

func TestTreeNode_InsertLevelOrder(t *testing.T) {
	tests := []struct {
		name    string
		root    int
		inserts []int
		want    []int
	}{
		{
			name:    "One level",
			root:    8,
			inserts: []int{},
			want:    []int{8},
		},
		{
			name:    "Two levels, imbalanced",
			root:    8,
			inserts: []int{3},
			want:    []int{8, 3},
		},
		{
			name:    "Two levels",
			root:    8,
			inserts: []int{3, 10},
			want:    []int{8, 3, 10},
		},
		{
			name:    "Three level, third level only left but incomplete",
			root:    8,
			inserts: []int{3, 10, 1},
			want:    []int{8, 3, 10, 1},
		},
		{
			name:    "Three level, third level only left",
			root:    8,
			inserts: []int{3, 10, 1, 6},
			want:    []int{8, 3, 10, 1, 6},
		},
		{
			name:    "Three level, third level right not complete",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14},
			want:    []int{8, 3, 10, 1, 6, 14},
		},
		{
			name:    "Three level",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14, 9},
			want:    []int{8, 3, 10, 1, 6, 14, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTreeNode(tt.root)

			for _, item := range tt.inserts {
				tr.InsertLevelOrder(item)
			}

			got := tr.LevelOrder()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.LevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_LevelOrder(t *testing.T) {
	tests := []struct {
		name    string
		root    int
		inserts []int
		want    []int
	}{
		{
			name:    "One level",
			root:    8,
			inserts: []int{},
			want:    []int{8},
		},
		{
			name:    "Two level",
			root:    8,
			inserts: []int{3, 10},
			want:    []int{8, 3, 10},
		},
		{
			name:    "Three level, third level only left but incomplete",
			root:    8,
			inserts: []int{3, 10, 1},
			want:    []int{8, 3, 10, 1},
		},
		{
			name:    "Three level, third level only left",
			root:    8,
			inserts: []int{3, 10, 1, 6},
			want:    []int{8, 3, 10, 1, 6},
		},
		{
			name:    "Three level, third level right not complete",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14},
			want:    []int{8, 3, 10, 1, 6, 14},
		},
		{
			name:    "Three level",
			root:    8,
			inserts: []int{3, 10, 1, 6, 14, 9},
			want:    []int{8, 3, 10, 1, 6, 14, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTreeNode(tt.root)

			for _, item := range tt.inserts {
				tr.InsertLevelOrder(item)
			}

			got := tr.LevelOrder()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.LevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_GetListOfDepths(t *testing.T) {
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

			got := tr.GetListOfDepths()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.GetListOfDepths() = %v, want %v", got, tt.want)
			}
		})
	}
}
