package treap

type Node struct {
	priority int
	val      int
	count    int
	l        *Node
	r        *Node
}

type Treap struct {
	Root *Node
	Seed int
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
		return a
	} else {
		b.l = merge(a, b.l)
		return b
	}
}

func split(key int, node *Node) (l *Node, r *Node) {
	if node == nil {
		return
	}
	if node.val < key {
		l = node
		l.r, r = split(key, node.r)
	} else {
		r = node
		l, r.l = split(key, node.l)
	}
	return
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

func (t *Treap) MakeNode(val int) *Node {
	return &Node{
		priority: t.Rand(),
		val:      val,
		count:    1,
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

func (t *Treap) Inorder(node *Node) {
	if node == nil {
		return
	}
	t.Inorder(node.l)
	t.Inorder(node.r)
}

func (t *Treap) Rand() int {
	t.Seed = (1103515245*t.Seed + 12345) % 2147483648
	return t.Seed
}
