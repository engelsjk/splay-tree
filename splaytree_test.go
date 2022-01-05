package splaytree

import (
	"fmt"
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

func TestSplay11(t *testing.T) {
	tree := New(less)
	values := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for _, v := range values {
		tree.Insert(v)
	}

	tree.Find(5.0)
	tree.Find(4.0)
	tree.Find(2.0)
	tree.Find(1.0)
	tree.Find(6.0)
	tree.Find(8.0)
	tree.Find(10.0)

	tree.PrintItems()
	// https://www.cs.umd.edu/class/fall2020/cmsc420-0201/Lects/lect10-splay.pdf
	// Fig. 4

	tree.Find(3.0)
	fmt.Println("splay 3")
	tree.PrintItems()
}

func TestSplay12(t *testing.T) {
	tree := New(less)

	values := []float64{6, 4, 3, 2, 5}

	for _, v := range values {
		tree.Insert(v)
	}

	tree.PrintItems()
	// http://www.btechsmartclass.com/data_structures/splay-trees.html

	tree.Find(3.0)
	fmt.Println("splay 3")
	tree.PrintItems()

	tree.Find(5.0)
	fmt.Println("splay 5")
	tree.PrintItems()

	tree.Find(2.0)
	fmt.Println("splay 2")
	tree.PrintItems()
}

func TestSplay21(t *testing.T) {
	tree := New(less)
	values := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for _, v := range values {
		tree.Insert(v)
	}

	tree.Find(2.0)
	tree.Find(6.0)
	tree.Find(8.0)
	tree.Find(10.0)

	tree.PrintItems()
	// https://www.cs.umd.edu/class/fall2020/cmsc420-0201/Lects/lect10-splay.pdf
	// Fig. 4

	tree.Find(3.0)
	fmt.Println("splay 3")
	tree.PrintItems()
}

func TestSplay22(t *testing.T) {
	tree := New(less)
	values := []float64{2, 3, 4, 5, 6}
	for _, v := range values {
		tree.Insert(v)
	}
	tree.Find(3.0)
	tree.Find(5.0)

	tree.PrintItems()
	// http://www.btechsmartclass.com/data_structures/splay-trees.html

	tree.Find(3.0)
	fmt.Println("splay 3")
	tree.PrintItems()

	tree.Find(5.0)
	fmt.Println("splay 5")
	tree.PrintItems()

	tree.Find(2.0)
	fmt.Println("splay 2")
	tree.PrintItems()
}

func TestPrintItems(t *testing.T) {
	tree := New(less)
	values := []float64{2, 12, 1, -6, 4, -8}
	for _, v := range values {
		tree.Insert(v)
	}
	tree.PrintItems()
}

// values := []float64{2, 3, 4, 5, 6}
// Perm(values, func(a []float64) {
// 	if a[4] != 5 {
// 		return
// 	}
// 	fmt.Printf("%+v\n", a)
// 	tree := New(less)
// 	for _, v := range values {
// 		tree.Insert(v)
// 	}
// 	tree.PrintItems()
// })

// Perm calls f with each permutation of a.
func Perm(a []float64, f func([]float64)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []float64, f func([]float64), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
