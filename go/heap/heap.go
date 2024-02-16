package heap

type Node struct {
	val int
}

type heap struct {
	arr  []Node
	less func(int, int) bool
	n    int
}

func makeNode(val int) Node {
	return Node{
		val: val,
	}
}

func (h *heap) Heapify(arr []int) {
	h.n = len(arr)
	h.arr = make([]Node, h.n)
	for i := 0; i < h.n; i++ {
		h.arr[i] = makeNode(arr[i])
	}
	h.less = func(i, j int) bool {
		return h.arr[i].val < h.arr[j].val
	}
	for i := h.n>>1 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *heap) siftDown(i int) {
	ci := i*2 + 1
	for ci < h.n {
		if ci+1 < h.n && h.less(ci+1, ci) {
			ci++
		}
		if h.less(ci, i) {
			h.arr[i], h.arr[ci] = h.arr[ci], h.arr[i]
			i = ci
			ci = i*2 + 1
		} else {
			break
		}
	}
}

func (h *heap) siftUp(i int) {
	var pi int
	for i > 0 {
		pi = (i - 1) >> 1
		if h.less(i, pi) {
			h.arr[i], h.arr[pi] = h.arr[pi], h.arr[i]
			i = pi
		} else {
			break
		}
	}
}

func (h *heap) Push(val int) {
	node := makeNode(val)
	h.arr = append(h.arr, node)
	h.n++
	h.siftUp(h.n - 1)

}

func (h *heap) Pop() Node {
	x := h.arr[0]
	h.arr[0] = h.arr[h.n-1]
	h.arr = h.arr[:h.n-1]
	h.n -= 1
	h.siftDown(0)
	return x
}

func (h *heap) Front() Node {
	return h.arr[0]
}

func (h *heap) Len() int {
	return h.n
}

func (h *heap) Less(less func(int, int) bool) {
	h.less = less
}
