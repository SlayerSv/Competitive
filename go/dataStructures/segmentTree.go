package dataStructures

type seg struct {
	tree []Node
	n    int
}

type Node struct {
	val   int
	count int
}

func (seg *seg) update(v1, v2 Node) Node {
	if v1.val < v2.val {
		return v1
	}
	if v1.val == v2.val {
		v2.count += v1.count
	}
	return v2
}

func makeNode(i int) Node {
	return Node{
		val:   i,
		count: 1,
	}
}

func (seg *seg) get(ql, qr, si, l, r int) Node {
	if qr <= l || ql >= r {
		return Node{int(1e10), 0}
	}
	if l >= ql && r <= qr {
		return seg.tree[si]
	}
	mid := (l + r) >> 1
	return seg.update(seg.get(ql, qr, si*2+1, l, mid), seg.get(ql, qr, si*2+2, mid, r))

}

func (seg *seg) set(i int, node Node, si, l, r int) {
	if r-l == 1 {
		seg.tree[si] = node
		return
	}
	mid := (l + r) >> 1
	if i < mid {
		seg.set(i, node, si*2+1, l, mid)
	} else {
		seg.set(i, node, si*2+2, mid, r)
	}
	seg.tree[si] = seg.update(seg.tree[si*2+1], seg.tree[si*2+2])
}

func (seg *seg) Get(ql, qr int) Node {
	return seg.get(ql, qr, 0, 0, seg.n)
}

func (seg *seg) Set(i, val int) {
	node := makeNode(val)
	seg.set(i, node, 0, 0, seg.n)
}

func (seg *seg) build(arr []int) {
	n := len(arr)
	l := 1
	for l < n {
		l *= 2
	}
	seg.tree = make([]Node, l*2-1)
	seg.n = l
	for i := 0; i < n; i++ {
		seg.Set(i, arr[i])
	}
}
