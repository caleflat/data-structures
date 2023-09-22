package datastructures

import (
	"testing"
)

func TestBTree(t *testing.T) {
	cases := []struct {
		name string
		data []int
	}{
		{
			name: "empty tree",
			data: []int{},
		},
		{
			name: "non-empty tree",
			data: []int{1, 2, 3, 4, 5},
		},
		{
			name: "tree with duplicates",
			data: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tree := NewBTree[int]()
			if !tree.IsEmpty() {
				t.Errorf("expected tree to be empty")
			}

			for _, v := range c.data {
				tree.Insert(v)
			}

			for _, v := range c.data {
				if !tree.Search(v) {
					t.Errorf("expected tree to contain %d", v)
				}
			}

			for _, v := range c.data {
				tree.Delete(v)
				if tree.Search(v) {
					t.Errorf("expected tree to not contain %d", v)
				}
			}

			if !tree.IsEmpty() {
				t.Errorf("expected tree to be empty")
			}
		})
	}
}
