package graphs

func topsort(graph [][]int) ([]int, bool) {
	n := len(graph) + 1
	indegrees := make([]int, n)
	for _, vertices := range graph {
		for _, v := range vertices {
			indegrees[v]++
		}
	}
	q := []int{}
	for i := 1; i < n; i++ {
		if indegrees[i] == 0 {
			q = append(q, i)
		}
	}
	path := []int{}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		path = append(path, node)
		for _, v := range graph[node] {
			indegrees[v]--
			if indegrees[v] == 0 {
				q = append(q, v)
			}
		}
	}
	hasCycle := len(path) != len(graph)
	return path, hasCycle
}
