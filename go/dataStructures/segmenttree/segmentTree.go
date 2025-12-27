package segmenttree

type seg struct {
    tree []Node
    // n is the number of leaves of a segtree
    n    int
}

type Node struct {
    val int
}

func (seg *seg) combine(v1, v2 Node) Node {
    node := makeNode(v1.val + v2.val)
    return node
}

func makeNode(i int) Node {
    return Node{
        val: i,
    }
}

func identity() Node {
    return makeNode(0)
}

// Node and Query cover interval [l, r)
func (seg *seg) Get(ql, qr int) Node {
    return seg.get(ql, qr, 1, 0, seg.n)
}

func (seg *seg) get(ql, qr, si, l, r int) Node {
    if qr <= l || r <= ql {
         return identity()
    }
    if ql <= l && r <= qr {
        return seg.tree[si]
    }
    mid := (l + r) >> 1
    return seg.combine(
        seg.get(ql, qr, si*2, l, mid),
        seg.get(ql, qr, si*2+1, mid, r),
    )
}

// Node and Query cover interval [l, r)
func (seg *seg) Set(ql, qr int, node Node) {
    seg.set(ql, qr, 1, 0, seg.n, node)
}

func (seg *seg) set(ql, qr, si, l, r int, node Node) {
    if r-l == 1 {
        seg.tree[si] = node
        return
    }
    mid := (l + r) >> 1
    if ql < mid {
        seg.set(ql, qr, si*2, l, mid, node)
    }
    if qr > mid {
        seg.set(ql, qr, si*2+1, mid, r, node)
    }
    seg.tree[si] = seg.combine(seg.tree[si*2], seg.tree[si*2+1])
}

func (seg *seg) build(arr []int) {
    n := 1
    for n < len(arr) {
        n <<= 1
    }
    seg.n = n
    seg.tree = make([]Node, 2*n)
    for i := 0; i < len(arr); i++ {
        seg.tree[n+i] = makeNode(arr[i])
    }
    for i := n - 1; i >= 1; i-- {
        seg.tree[i] = seg.combine(seg.tree[i*2], seg.tree[i*2+1])
    }
}