package main

import (
	"fmt"

	splaytree "github.com/engelsjk/splay-tree"
)

type Item struct {
	Key int
}

func byKeys(a, b interface{}) int {
	i1, i2 := a.(*Item), b.(*Item)
	if i1.Key == i2.Key {
		return 0
	} else if i1.Key < i2.Key {
		return -1
	} else {
		return 1
	}
}

func main() {

	t := splaytree.New(byKeys)
	t.Insert(&Item{5})
	t.Insert(&Item{-10})
	t.Insert(&Item{0})
	t.Insert(&Item{33})
	t.Insert(&Item{2})
	ni := t.Insert(&Item{4})
	t.Add(&Item{10})

	t.Find(&Item{4})
	t.Remove(&Item{4})

	for _, node := range t.Nodes() {
		fmt.Printf("(%p) %+v: %+v\n", node, node, node.Item())
	}

	fmt.Printf("ni: (%p) %+v\n", ni, ni)
	ni_next := t.Next(&splaytree.Node{})
	fmt.Printf("n10_next: (%p) %+v\n", ni_next, ni_next)
	ni_prev := t.Prev(ni)
	fmt.Printf("n10_prev: (%p) %+v\n", ni_prev, ni_prev)

	fmt.Printf("size: %d\n", t.Size())
	fmt.Printf("min: %+v\n", t.Min())
	fmt.Printf("max: %+v\n", t.Max())

	t.Remove(&Item{0})
	fmt.Printf("size: %d\n", t.Size())

	t.Remove(&Item{33})
	fmt.Printf("size: %d\n", t.Size())
	fmt.Printf("min: %+v\n", t.Min())
	fmt.Printf("max: %+v\n", t.Max())
}
