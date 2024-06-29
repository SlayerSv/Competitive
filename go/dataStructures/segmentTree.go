package dataStructures

type seg struct {
	tree []Node
	n    int
}

type Node struct {
	val int
}

func (seg *seg) update(v1, v2 Node) Node {
	v1.val += v2.val
	return v1
}

func makeNode(i int) Node {
	return Node{
		val: i,
	}
}

func (seg *seg) get(ql, qr, si, l, r int) Node {
	if l >= ql && r <= qr {
		return seg.tree[si]
	}
	mid := (l + r) >> 1
	if qr <= mid {
		return seg.get(ql, qr, si*2+1, l, mid)
	} else if ql >= mid {
		return seg.get(ql, qr, si*2+2, mid, r)
	} else {
		return seg.update(seg.get(ql, qr, si*2+1, l, mid), seg.get(ql, qr, si*2+2, mid, r))
	}
}

func (seg *seg) set(ql, qr, val, si, l, r int) {
	if r-l == 1 {
		node := makeNode(val)
		seg.tree[si] = seg.update(seg.tree[si], node)
		return
	}
	mid := (l + r) >> 1
	if ql < mid {
		seg.set(ql, qr, val, si*2+1, l, mid)
	}
	if qr > mid {
		seg.set(ql, qr, val, si*2+2, mid, r)
	}
	seg.tree[si] = seg.update(seg.tree[si*2+1], seg.tree[si*2+2])
}

func (seg *seg) Get(ql, qr int) Node {
	return seg.get(ql, qr, 0, 0, seg.n)
}

func (seg *seg) Set(ql, qr int, val int) {
	seg.set(ql, qr, val, 0, 0, seg.n)
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
		seg.Set(i, i+1, arr[i])
	}
}
