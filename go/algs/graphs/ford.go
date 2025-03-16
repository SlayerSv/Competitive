package graphs

type Edge struct {
	from int
	to   int
	dist int
}

// negCycleVert is -1 if there is no cycle, else it contains any vertex belonging
// to the negative cycle.
func Ford(start, n int, edges []Edge) (dist []int, from []int, negCycle []int) {
	dist = make([]int, n+1)
	from = make([]int, n+1)
	for i := 0; i < len(dist); i++ {
		dist[i] = 1e18
		from[i] = -1
	}
	dist[start] = 0
	negCycleVert := -1
	for i := 0; i < n; i++ {
		negCycleVert = -1
		for _, edge := range edges {
			if dist[edge.from] == 1e18 {
				continue
			}
			if dist[edge.from]+edge.dist < dist[edge.to] {
				dist[edge.to] = dist[edge.from] + edge.dist
				from[edge.to] = edge.from
				negCycleVert = edge.to
			}
		}
		if negCycleVert == -1 {
			break
		}
	}
	if negCycleVert != -1 {
		for i := 0; i < n; i++ {
			negCycleVert = from[negCycleVert]
		}
		u := negCycleVert
		for u != negCycleVert || len(negCycle) < 1 {
			negCycle = append(negCycle, u)
			u = from[u]
		}
	}
	return
}

type V struct {
	v    int
	dist int
}

func FordFA(start int, graph [][]V) (dist []int, from []int, hasCycle bool) {
	dist = make([]int, len(graph)+1)
	from = make([]int, len(graph)+1)
	for i := 0; i < len(dist); i++ {
		dist[i] = 1e18
		from[i] = -1
	}
	dist[start] = 0
	q := make([]int, 0, len(graph))
	inq := make([]bool, len(graph)+1)
	q = append(q, start)
	inq[start] = true
	relaxed := make([]int, len(graph)+1)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		inq[v] = false
		for _, to := range graph[v] {
			d := dist[v] + to.dist
			if d < dist[to.v] {
				dist[to.v] = d
				relaxed[to.v]++
				if relaxed[to.v] > len(graph) {
					hasCycle = true
					break
				}
				if !inq[to.v] {
					q = append(q, to.v)
					inq[to.v] = true
				}
			}
		}
		if hasCycle {
			break
		}
	}
	return
}
