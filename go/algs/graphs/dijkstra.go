package graphs

func Dijsktra(graph [][][2]int, node int) ([]int, []int) {
	n := len(graph)
	dists := make([]int, n)
	from := make([]int, n)
	const INF int = 1e18
	for i := 0; i < n; i++ {
		dists[i] = INF
		from[i] = -1
	}
	dists[node] = 0
	pq := heap{}
	pq.Push(makeNode(node, 0))
	for pq.Len() > 0 {
		node := pq.Pop()
		if node.dist > dists[node.v] {
			continue
		}
		for _, neigh := range graph[node.v] {
			dist := dists[node.v] + neigh[1]
			if dist >= dists[neigh[0]] {
				continue
			}
            dists[neigh[0]] = dist
			from[neigh[0]] = node.v
			pq.Push(makeNode(neigh[0], dist))
		}
	}
	return dists, from
}

type Node struct {
	v    int
	dist int
}

type heap struct {
	arr []Node
	n   int
}

func makeNode(v, dist int) Node {
	return Node{
		v:    v,
		dist: dist,
	}
}

func (h *heap) Less(i, j int) bool {
	return h.arr[i].dist < h.arr[j].dist
}

func (h *heap) Heapify(arr []int) {
	h.n = len(arr)
	h.arr = make([]Node, h.n)
	for i := 0; i < h.n; i++ {
		h.arr[i] = makeNode(arr[i], 0)
	}
	for i := h.n>>1 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *heap) siftDown(i int) {
	ci := i*2 + 1
	for ci < h.n {
		if ci+1 < h.n && h.Less(ci+1, ci) {
			ci++
		}
		if h.Less(ci, i) {
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
		if h.Less(i, pi) {
			h.arr[i], h.arr[pi] = h.arr[pi], h.arr[i]
			i = pi
		} else {
			break
		}
	}
}

func (h *heap) Push(node Node) {
	h.arr = append(h.arr, node)
	h.n++
	h.siftUp(h.n - 1)
}

func (h *heap) Pop() Node {
	if h.n == 0 {
		return Node{}
	}
	x := h.arr[0]
	h.arr[0] = h.arr[h.n-1]
	h.arr = h.arr[:h.n-1]
	h.n -= 1
	h.siftDown(0)
	return x
}

func (h *heap) Front() Node {
	if h.n == 0 {
		return Node{}
	}
	return h.arr[0]
}

func (h *heap) Len() int {
	return h.n
}
