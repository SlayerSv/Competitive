package treap

import (
    "math/rand"
)

type Node struct {
	priority int
	val      int
	count    int
	l        *Node
	r        *Node
	size     int // size of the subtree
}

type Treap struct {
	Root *Node
	Size int
}

func NewTreap() *Treap {
    return &Treap{}
}

func merge(a, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.priority > b.priority {
		a.r = merge(a.r, b)
		recalcSize(a)
		return a
	} else {
		b.l = merge(a, b.l)
		recalcSize(b)
		return b
	}
}

// l = (lowest...key-1)  r = (key...highest)
func split(key int, node *Node) (l *Node, r *Node) {
	if node == nil {
		return
	}
	if key <= node.val {
        r = node
		l, r.l = split(key, node.l)
	} else {
		l = node
		l.r, r = split(key, node.r)
	}
	recalcSize(node)
	return
}

func recalcSize(node *Node) {
	node.size = node.count
	if node.l != nil {
		node.size += node.l.size
	}
	if node.r != nil {
		node.size += node.r.size
	}
}

func (t *Treap) Add(val int) {
	var l, equal, r *Node
	l, r = split(val, t.Root)
	equal, r = split(val+1, r)
	if equal != nil {
		equal.count++
	} else {
		equal = t.MakeNode(val)
	}
	t.Root = merge(l, merge(equal, r))
	t.Size++
}

func (t *Treap) Delete(val int) {
	var l, equal, r *Node
	l, r = split(val, t.Root)
	equal, r = split(val+1, r)
	if equal != nil && equal.count > 1 {
		equal.count--
		t.Root = merge(l, merge(equal, r))
		return
	}
	t.Root = merge(l, r)
	t.Size--
}

func (t *Treap) Has(val int) bool {
	node := t.Root
	for node != nil {
		if node.val == val {
			return true
		}
		if val < node.val {
			node = node.l
		} else {
			node = node.r
		}
	}
	return false
}

// 0-based index
func (t *Treap) GetByIndex(i int) int {
    val := -1
	node := t.Root
    i++ // turn to 1-based index for more simple logic
	for {
        if node == nil {
            break
        }
		if node.l != nil {
			if i <= node.l.size {
				node = node.l
				continue
			}
			i -= node.l.size
		}
		if i == 1 {
			val = node.val
			break
		}
		node = node.r
		i--
	}
	return val
}

func (t *Treap) MakeNode(val int) *Node {
	return &Node{
		priority: rand.Int(),
		val:      val,
		count:    1,
		size:     1,
	}
}

func (t *Treap) Lower(val int) int {
	l, r := split(val+1, t.Root)
	if l == nil {
		return -1
	}
	node := l
	for node.r != nil {
		node = node.r
	}
	merge(l, r)
	return node.val
}

func (t *Treap) Upper(val int) int {
	l, r := split(val, t.Root)
	if r == nil {
		return -1
	}
	node := r
	for node.l != nil {
		node = node.l
	}
	merge(l, r)
	return node.val
}

func (t *Treap) Min() int {
	tmp := t.Root
	for tmp != nil && tmp.l != nil {
		tmp = tmp.l
	}
	if tmp != nil {
		return tmp.val
	}
	return -1
}

func (t *Treap) Max() int {
	tmp := t.Root
	for tmp != nil && tmp.r != nil {
		tmp = tmp.r
	}
	if tmp != nil {
		return tmp.val
	}
	return -1
}

func (t *Treap) Inorder(node *Node) {
	if node == nil {
		return
	}
	t.Inorder(node.l)
	t.Inorder(node.r)
}
