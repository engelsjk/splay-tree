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
