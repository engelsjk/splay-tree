package splaytree

// Follows "An implementation of top-down splaying"
// by D. Sleator <sleator@cs.cmu.edu> March 1992

type SplayTree struct {
	root       *Node
	size       int
	comparator func(a, b interface{}) int
}

type Node struct {
	item        interface{}
	left, right *Node
}

func New(less func(a, b interface{}) int) *SplayTree {
	return &SplayTree{
		comparator: less,
	}
}

// Insert a key, allows duplicates
func (tr *SplayTree) Insert(item interface{}) *Node {
	tr.size++
	tr.root = insert(item, tr.root, tr.comparator)
	return tr.root
}

// Add a key, if it is not present in the tree
func (tr *SplayTree) Add(item interface{}) *Node {
	n := new(Node)
	if tr.root == nil {
		n.left = nil
		n.right = nil
		tr.size++
		tr.root = n
	}
	t := splay(item, tr.root, tr.comparator)
	cmp := tr.comparator(item, t.item)
	if cmp == 0 {
		tr.root = t
	} else {
		if cmp < 0 {
			n.left = t.left
			n.right = t
			t.left = nil
		} else if cmp > 0 {
			n.right = t.right
			n.left = t
			t.right = nil
		}
		tr.size++
		tr.root = n
	}
	return tr.root
}

// Remove i from the tree if it's there
func (tr *SplayTree) Remove(item interface{}) {
	tr.root = tr.remove(item, tr.root, tr.comparator)
}

func (tr *SplayTree) remove(
	i interface{},
	t *Node,
	comparator func(a, b interface{}) int,
) *Node {
	if t == nil {
		return nil
	}
	t = splay(i, t, comparator)
	cmp := comparator(i, t.item)
	if cmp == 0 {
		var x *Node
		if t.left == nil {
			x = t.right
		} else {
			x = splay(i, t.left, comparator)
			x.right = t.right
		}
		tr.size--
		return x
	}
	return t
}

// Pop removes and returns the node with smallest key
func (tr *SplayTree) Pop() interface{} {
	n := tr.root
	if n == nil {
		return nil
	}
	if n.left != nil { // check: for loop?
		n = n.left
	}
	tr.root = splay(n.item, tr.root, tr.comparator)
	tr.root = tr.remove(n.item, tr.root, tr.comparator)
	return n.item
}

// FindStatic finds without splaying
// func (tr *SplayTree) FindStatic() {}

func (tr *SplayTree) Find(item interface{}) *Node {
	if tr.root != nil {
		tr.root = splay(item, tr.root, tr.comparator)
		if tr.comparator(item, tr.root.item) != 0 {
			return nil
		}
	}
	return tr.root
}

func (tr *SplayTree) Contains(item interface{}) bool {
	current := tr.root
	for current != nil {
		cmp := tr.comparator(item, current.Item())
		if cmp == 0 {
			return true
		} else if cmp < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}
	return false
}

// func (tr *SplayTree) ForEach() {}
// func (tr *SplayTree) Range() {}
// func (tr *SplayTree) Keys() {}
// func (tr *SplayTree) Values() {}

func (tr *SplayTree) Min() interface{} {
	if tr.root == nil {
		return nil
	}
	return tr.MinNode(tr.root).item
}

func (tr *SplayTree) Max() interface{} {
	if tr.root == nil {
		return nil
	}
	return tr.MaxNode(tr.root).item
}

func (tr *SplayTree) MinNode(t *Node) *Node {
	if t == nil {
		t = tr.root
	}
	for t.left != nil {
		t = t.left
	}
	return t
}

func (tr *SplayTree) MaxNode(t *Node) *Node {
	if t == nil {
		t = tr.root
	}
	for t.right != nil {
		t = t.right
	}
	return t
}

// At returns node at given index
// func (tr *SplayTree) At() {}

func (tr *SplayTree) Next(d *Node) *Node {
	root := tr.root
	var successor *Node
	if d.right != nil {
		successor = d.right
		for successor.left != nil {
			successor = successor.left
		}
		return successor
	}
	for root != nil {
		cmp := tr.comparator(d.item, root.item)
		if cmp == 0 {
			break
		} else if cmp < 0 {
			successor = root
			root = root.left
		} else {
			root = root.right
		}
	}
	return successor
}

func (tr *SplayTree) Prev(d *Node) *Node {
	root := tr.root
	var predecessor *Node
	if d.left != nil {
		predecessor = d.left
		for predecessor.right != nil {
			predecessor = predecessor.right
		}
		return predecessor
	}

	for root != nil {
		cmp := tr.comparator(d.item, root.item)
		if cmp == 0 {
			break
		} else if cmp < 0 {
			root = root.left
		} else {
			predecessor = root
			root = root.right
		}
	}
	return predecessor
}

// func (tr *SplayTree) Clear() {}
// func (tr *SplayTree) ToList() {}
// func (tr *SplayTree) Load() {}

func (tr *SplayTree) IsEmpty() bool {
	return tr.root == nil
}

func (tr *SplayTree) Size() int {
	return tr.size
}

// func (tr *SplayTree) Root() {}
// func (tr *SplayTree) ForEach() {}
// func (tr *SplayTree) ToString() {}
// func (tr *SplayTree) Update() {}
// func (tr *SplayTree) Split() {}

func (n *Node) Item() interface{} {
	return n.item
}

func insert(
	i interface{},
	t *Node,
	comparator func(a, b interface{}) int,
) *Node {

	n := &Node{item: i}

	if t == nil {
		n.left, n.right = nil, nil
		return n
	}

	t = splay(i, t, comparator)
	cmp := comparator(i, t.item)
	if cmp < 0 {
		n.left = t.left
		n.right = t
		t.left = nil
	} else if cmp >= 0 {
		n.right = t.right
		n.left = t
		t.right = nil
	}
	return n
}

// Simple top down splay, not requiring i to be in the tree t.
func splay(
	i interface{},
	t *Node,
	comparator func(a, b interface{}) int,
) *Node {
	n := new(Node)
	l := n
	r := n
	for {
		cmp := comparator(i, t.item)
		if cmp < 0 {
			if t.left == nil {
				break
			}
			if comparator(i, t.left.item) < 0 {
				y := t.left // rotate right
				t.left = y.right
				y.right = t
				t = y
				if t.left == nil {
					break
				}
			}
			r.left = t // link right
			r = t
			t = t.left
		} else if cmp > 0 {
			if t.right == nil {
				break
			}
			if comparator(i, t.right.item) > 0 {
				y := t.right // rotate left
				t.right = y.left
				y.left = t
				t = y
				if t.right == nil {
					break
				}
			}
			l.right = t // link left
			l = t
			t = t.right
		} else {
			break
		}
	}
	// assemble
	l.right = t.left
	r.left = t.right
	t.left = n.right
	t.right = n.left
	return t
}
