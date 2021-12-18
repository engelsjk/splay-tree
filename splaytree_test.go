package splaytree

import (
	"testing"
)

func expect(t testing.TB, what bool) {
	t.Helper()
	if !what {
		t.Fatal("expection failure")
	}
}

func less(a, b interface{}) int {
	af := a.(float64)
	bf := b.(float64)
	if af > bf {
		return 1
	}
	if af < bf {
		return -1
	}
	return 0
}

func TestDuplicate(t *testing.T) {

	var tree *SplayTree
	var values []float64
	var size int

	// TODO: should allow inserting of duplicate key
	// TODO: should allow multiple duplicate keys in a row

	// TODO: should remove from a tree with duplicate keys correctly
	tree = New(less)
	values = []float64{2, 12, 1, 1, -6, 1, 1}
	for _, v := range values {
		tree.Insert(v)
	}
	size = tree.Size()
	for i := 0; i < 4; i++ {
		tree.Remove(1.0)
		if i < 3 {
			expect(t, tree.Contains(1.0))
		}
		size--
		expect(t, tree.Size() == size)
	}

	// should remove from a tree with multiple duplicate keys correctly
	tree = New(less)
	values = []float64{2, 12, 1, 1, -6, 1, 1, 2, 0, 2}
	for _, v := range values {
		tree.Insert(v)
	}
	size = tree.Size()
	for !tree.IsEmpty() {
		tree.Pop()
		size--
		expect(t, tree.Size() == size)
	}

	// TODO: should disallow duplicates if noDuplicates is set

	// should add only if the key is not there
	tree = New(less)
	tree.Insert(1.0)
	tree.Insert(2.0)
	tree.Insert(3.0)
	size = tree.Size()
	tree.Add(1.0)
	expect(t, tree.Size() == size)
}
