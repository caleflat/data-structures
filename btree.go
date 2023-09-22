package datastructures

import (
	"golang.org/x/exp/constraints"
)

// BTree is a binary tree data structure
// It is a tree data structure in which each node has at most two children,
// which are referred to as the left child and the right child.
// A recursive definition using just set theory notions is that a (non-empty) binary tree is a tuple (L, S, R),
// where L and R are binary trees or the empty set and S is a singleton set containing the root.
type BTree[T constraints.Ordered] struct {
	Root *BTreeNode[T]
}

// BTreeNode is a node in a binary tree
type BTreeNode[T constraints.Ordered] struct {
	Left  *BTreeNode[T]
	Right *BTreeNode[T]
	Value T
}

// NewBTree creates a new binary tree
func NewBTree[T constraints.Ordered]() *BTree[T] {
	return &BTree[T]{}
}

// Insert inserts a value into the binary tree
func (t *BTree[T]) Insert(value T) {
	if t.Root == nil {
		t.Root = &BTreeNode[T]{Value: value}
		return
	}

	t.Root.Insert(value)
}

// Insert inserts a value into the binary tree
func (n *BTreeNode[T]) Insert(value T) {
	if value == n.Value {
		return
	}

	if value < n.Value {
		if n.Left == nil {
			n.Left = &BTreeNode[T]{Value: value}
			return
		}

		n.Left.Insert(value)
		return
	}

	if n.Right == nil {
		n.Right = &BTreeNode[T]{Value: value}
		return
	}

	n.Right.Insert(value)
}

// Search searches for a value in the binary tree
func (t *BTree[T]) Search(value T) bool {
	if t.Root == nil {
		return false
	}

	return t.Root.Search(value)
}

// Search searches for a value in the binary tree
func (n *BTreeNode[T]) Search(value T) bool {
	if value == n.Value {
		return true
	}

	if value < n.Value {
		if n.Left == nil {
			return false
		}

		return n.Left.Search(value)
	}

	if n.Right == nil {
		return false
	}

	return n.Right.Search(value)
}

// Delete deletes a value from the binary tree
func (t *BTree[T]) Delete(value T) {
	if t.Root == nil {
		return
	}

	t.Root = t.Root.Delete(value)
}

// Delete deletes a value from the binary tree
func (n *BTreeNode[T]) Delete(value T) *BTreeNode[T] {
	if value < n.Value {
		if n.Left == nil {
			return n
		}

		n.Left = n.Left.Delete(value)
		return n
	}

	if value > n.Value {
		if n.Right == nil {
			return n
		}

		n.Right = n.Right.Delete(value)
		return n
	}

	if n.Left == nil {
		return n.Right
	}

	if n.Right == nil {
		return n.Left
	}

	min := n.Right
	for min.Left != nil {
		min = min.Left
	}

	n.Value = min.Value
	n.Right = n.Right.Delete(min.Value)
	return n
}

// Traverse traverses the binary tree
func (t *BTree[T]) Traverse(f func(T)) {
	if t.Root == nil {
		return
	}

	t.Root.Traverse(f)
}

// Traverse traverses the binary tree
func (n *BTreeNode[T]) Traverse(f func(T)) {
	if n.Left != nil {
		n.Left.Traverse(f)
	}

	f(n.Value)

	if n.Right != nil {
		n.Right.Traverse(f)
	}
}

// Size returns the size of the binary tree
func (t *BTree[T]) Size() int {
	if t.Root == nil {
		return 0
	}

	return t.Root.Size()
}

// Size returns the size of the binary tree
func (n *BTreeNode[T]) Size() int {
	size := 1

	if n.Left != nil {
		size += n.Left.Size()
	}

	if n.Right != nil {
		size += n.Right.Size()
	}

	return size
}

// Height returns the height of the binary tree
func (t *BTree[T]) Height() int {
	if t.Root == nil {
		return 0
	}

	return t.Root.Height()
}

// Height returns the height of the binary tree
func (n *BTreeNode[T]) Height() int {
	leftHeight := 0
	if n.Left != nil {
		leftHeight = n.Left.Height()
	}

	rightHeight := 0
	if n.Right != nil {
		rightHeight = n.Right.Height()
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

// IsBalanced returns true if the binary tree is balanced
func (t *BTree[T]) IsBalanced() bool {
	if t.Root == nil {
		return true
	}

	return t.Root.IsBalanced()
}

// IsBalanced returns true if the binary tree is balanced
func (n *BTreeNode[T]) IsBalanced() bool {
	leftHeight := 0
	if n.Left != nil {
		leftHeight = n.Left.Height()
	}

	rightHeight := 0
	if n.Right != nil {
		rightHeight = n.Right.Height()
	}

	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		return false
	}

	if n.Left != nil && !n.Left.IsBalanced() {
		return false
	}

	if n.Right != nil && !n.Right.IsBalanced() {
		return false
	}

	return true
}

// IsEmpty returns true if the binary tree is empty
func (t *BTree[T]) IsEmpty() bool {
	return t.Root == nil
}

// IsEmpty returns true if the binary tree is empty
func (n *BTreeNode[T]) IsEmpty() bool {
	return n == nil
}
