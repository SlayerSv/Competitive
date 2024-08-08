package segmenttree

type lst struct {
	tree []lstNode
	n    int
}

type lstNode struct {
	val int
	op  bool
}

func modify(v1, v2 lstNode) lstNode {
	v1.val += v2.val
	return v1
}

func makelstNode(i int) lstNode {
	return lstNode{
		val: i,
	}

}

func (lst *lst) propagate(si int) {
	lst.tree[si*2+1] = modify(lst.tree[si*2+1], lst.tree[si])
	lst.tree[si*2+1].op = true
	lst.tree[si*2+2] = modify(lst.tree[si*2+2], lst.tree[si])
	lst.tree[si*2+2].op = true
	lst.tree[si].op = false
	lst.tree[si].val = 0
}

func (lst *lst) get(i, si, l, r int) lstNode {
	if r-l > 1 && lst.tree[si].op {
		lst.propagate(si)
	}
	if r-l == 1 {
		return lst.tree[si]
	}
	mid := (l + r) >> 1
	if i < mid {
		return lst.get(i, si*2+1, l, mid)
	} else {
		return lst.get(i, si*2+2, mid, r)
	}
}

func (lst *lst) set(ql, qr int, node lstNode, si, l, r int) {
	if l >= ql && r <= qr {
		lst.tree[si] = modify(lst.tree[si], node)
		lst.tree[si].op = true
		return
	}
	mid := (l + r) >> 1
	if ql < mid {
		lst.set(ql, qr, node, si*2+1, l, mid)
	}
	if qr > mid {
		lst.set(ql, qr, node, si*2+2, mid, r)
	}
}

func (lst *lst) Get(i int) lstNode {
	return lst.get(i, 0, 0, lst.n)
}

func (lst *lst) Set(ql, qr, val int) {
	node := makelstNode(val)
	lst.set(ql, qr, node, 0, 0, lst.n)
}

func (lst *lst) build(n int) {
	l := 1
	for l < n {
		l *= 2
	}
	lst.tree = make([]lstNode, l*2-1)
	lst.n = l
	// for i := 0; i < n; i++ {
	// 	lst.Set(0, lst.n, arr[i])
	// }
}
