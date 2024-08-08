package graphs

func mst(graph [][][2]int) {
	n := len(graph)
	connected := make([]bool, n)
	q := heap{}
	q.Push(makeNode(0, 0))
	var node Node
	for n > 0 {
		node = q.Pop()
		if connected[node.v] {
			continue
		}
		connected[node.v] = true
		n--
		for _, neigh := range graph[node.v] {
			to, dist := neigh[0], neigh[1]
			if connected[to] {
				continue
			}
			q.Push(makeNode(to, dist))
		}
	}
}
