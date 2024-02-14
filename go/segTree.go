package main

type seg struct {
	tree []Node
	n    int
}

type Node struct {
	val int
}

func (seg *seg) update(v1, v2 int) int {
	return v1 + v2
}

func (seg *seg) get(ql, qr, si, l, r int) int {
	if qr <= l || ql >= r {
		return 0
	}
	if l >= ql && r <= qr {
		return seg.tree[si].val
	}
	mid := (l + r) >> 1
	s1 := seg.get(ql, qr, si*2+1, l, mid)
	s2 := seg.get(ql, qr, si*2+2, mid, r)
	return seg.update(s1, s2)

}

func (seg *seg) set(i, val, si, l, r int) {
	if r-l == 1 {
		seg.tree[si].val = val
		return
	}

	mid := (l + r) >> 1
	if i < mid {
		seg.set(i, val, si*2+1, l, mid)
	} else {
		seg.set(i, val, si*2+2, mid, r)
	}
	seg.tree[si].val = seg.update(seg.tree[si*2+1].val, seg.tree[si*2+2].val)
}

func (seg *seg) Get(ql, qr int) int {
	return seg.get(ql, qr, 0, 0, seg.n)
}

func (seg *seg) Set(i, val int) {
	seg.set(i, val, 0, 0, seg.n)
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
