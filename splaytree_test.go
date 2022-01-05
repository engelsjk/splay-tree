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

func equal(a []interface{}, b []float64) bool {
	for i := 0; i < len(a); i++ {
		v, ok := a[i].(float64)
		if !ok {
			return false
		}
		if b[i] != v {
			return false
		}
	}
	return true
}

func TestInsert(t *testing.T) {
	var tree *SplayTree

	// should return the size of the tree
	tree = New(less)
	tree.Insert(1.0)
	tree.Insert(2.0)
	tree.Insert(3.0)
	tree.Insert(4.0)
	tree.Insert(5.0)
	expect(t, tree.Size() == 5)

	// should return the pointer
	tree = New(less)
	n1 := tree.Insert(1.0)
	n2 := tree.Insert(2.0)
	n3 := tree.Insert(3.0)
	expect(t, n1.Item().(float64) == 1.0)
	expect(t, n2.Item().(float64) == 2.0)
	expect(t, n3.Item().(float64) == 3.0)
}

func TestDuplicate(t *testing.T) {

	var tree *SplayTree
	var values []float64
	var size int

	// should allow inserting of duplicate key
	tree = New(less)
	values = []float64{2, 12, 1, -6, 1}
	for _, v := range values {
		tree.Insert(v)
	}
	expect(t, equal(tree.Items(), []float64{-6, 1, 1, 2, 12}))
	expect(t, tree.Size() == 5)

	// should allow multiple duplicate keys in a row
	tree = New(less)
	values = []float64{2, 12, 1, 1, -6, 2, 1, 1, 13}
	for _, v := range values {
		tree.Insert(v)
	}
	expect(t, equal(tree.Items(), []float64{-6, 1, 1, 1, 1, 2, 2, 12, 13}))
	expect(t, tree.Size() == 9)

	// should remove from a tree with duplicate keys correctly
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

	// should disallow duplicates if noDuplicates is set
	tree = New(less)
	values = []float64{2, 12, 1, -6, 1}
	for _, v := range values {
		tree.Add(v)
	}
	expect(t, equal(tree.Items(), []float64{-6, 1, 2, 12}))
	expect(t, tree.Size() == 4)

	// should add only if the key is not there
	tree = New(less)
	tree.Insert(1.0)
	tree.Insert(2.0)
	tree.Insert(3.0)
	size = tree.Size()
	tree.Add(1.0)
	expect(t, tree.Size() == size)
}

func TestPrev(t *testing.T) {
	tree := New(less)
	values := []float64{2, 12, 1, -6, 4, -8}
	for _, v := range values {
		tree.Insert(v)
	}

	node := tree.Find(4.0)
	prevNode := tree.Prev(node)

	expect(t, prevNode.Item().(float64) == 2.0)
}

func TestNext(t *testing.T) {
	tree := New(less)
	values := []float64{2, 12, 1, -6, 4, -8}
	for _, v := range values {
		tree.Insert(v)
	}

	node := tree.Find(4.0)
	prevNode := tree.Next(node)

	expect(t, prevNode.Item().(float64) == 12.0)
}

func TestStringify(t *testing.T) {
	tree := New(less)
	values := []float64{2, 12, 1, -6, 4, -8}
	for _, v := range values {
		tree.Insert(v)
	}
	tree.String()
}
