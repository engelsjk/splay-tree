package splaytree

type SplayTree struct {
	root       *node
	size       int
	comparator func(a, b interface{}) int
}

type node struct {
	item        interface{}
	left, right *node
}

func New(less func(a, b interface{}) int) *SplayTree {
	return &SplayTree{
		comparator: less,
	}
}

func (tr *SplayTree) Insert(item interface{}) *node {
	tr.size++
	tr.root = insert(item, tr.root, tr.comparator)
	return tr.root
}

func (tr *SplayTree) Add(item interface{}) *node {
	n := new(node)
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

func (tr *SplayTree) Remove(item interface{}) {
	tr.root = tr.remove(item, tr.root, tr.comparator)
}

func (tr *SplayTree) remove(
	i interface{},
	t *node,
	comparator func(a, b interface{}) int,
) *node {
	if t == nil {
		return nil
	}
	t = splay(i, t, comparator)
	cmp := comparator(i, t.item)
	if cmp == 0 {
		var x *node
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

// func (tr *SplayTree) FindStatic() {}

func (tr *SplayTree) Find(item interface{}) *node {
	if tr.root != nil {
		tr.root = splay(item, tr.root, tr.comparator)
		if tr.comparator(item, tr.root.item) != 0 {
			return nil
		}
	}
	return tr.root
}

// func (tr *SplayTree) Contains() {}
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

func (tr *SplayTree) MinNode(t *node) *node {
	if t == nil {
		t = tr.root
	}
	for t.left != nil {
		t = t.left
	}
	return t
}

func (tr *SplayTree) MaxNode(t *node) *node {
	if t == nil {
		t = tr.root
	}
	for t.right != nil {
		t = t.right
	}
	return t
}

// func (tr *SplayTree) At() {}

func (tr *SplayTree) Next(d *node) *node {
	root := tr.root
	var successor *node
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

func (tr *SplayTree) Prev(d *node) *node {
	root := tr.root
	var predecessor *node
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
// func (tr *SplayTree) IsEmpty() {}

func (tr *SplayTree) Size() int {
	return tr.size
}

// func (tr *SplayTree) Root() {}
// func (tr *SplayTree) ForEach() {}
// func (tr *SplayTree) ToString() {}
// func (tr *SplayTree) Update() {}
// func (tr *SplayTree) Split() {}

func (n *node) Item() interface{} {
	return n.item
}

func insert(
	i interface{},
	t *node,
	comparator func(a, b interface{}) int,
) *node {

	n := &node{item: i}

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

func splay(
	i interface{},
	t *node,
	comparator func(a, b interface{}) int,
) *node {
	n := new(node)
	l := n
	r := n
	for {
		cmp := comparator(i, t.item)
		if cmp < 0 {
			if t.left == nil {
				break
			}
			if comparator(i, t.left.item) < 0 {
				y := t.left
				t.left = y.right
				y.right = t
				t = y
				if t.left == nil {
					break
				}
			}
			r.left = t
			r = t
			t = t.left
		} else if cmp > 0 {
			if t.right == nil {
				break
			}
			if comparator(i, t.right.item) > 0 {
				y := t.right
				t.right = y.left
				y.left = t
				t = y
				if t.right == nil {
					break
				}
			}
			l.right = t
			l = t
			t = t.right
		} else {
			break
		}
	}
	l.right = t.left
	r.left = t.right
	t.left = n.right
	t.right = n.left
	return t
}
