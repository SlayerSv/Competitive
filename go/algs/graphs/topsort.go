package graphs

func kahn(graph [][]int) ([]int, bool) {
	n := len(graph) + 1 // if vertices start with 0 then n := len(graph)
	indegrees := make([]int, n)
	for _, from := range graph {
		for _, to := range from {
			indegrees[to]++
		}
	}
	q := make([]int, 0, n)
	for i := 1; i < n; i++ {
		if indegrees[i] == 0 {
			q = append(q, i)
		}
	}
	path := make([]int, 0, n)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		path = append(path, node)
		for _, to := range graph[node] {
			indegrees[to]--
			if indegrees[to] == 0 {
				q = append(q, to)
			}
		}
	}
	hasCycle := len(path) != len(graph)
	return path, hasCycle
}
