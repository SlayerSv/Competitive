package graphs

// graph must be an adjecancy matrix filled with INF values
// if there are no edge between 2 vertices.
// function will modify the graph in place.
func floyd(graph [][]int) {
	n := len(graph)
	const INF int = 1e18
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if graph[i][k] < INF && graph[k][j] < INF {
					graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
				}
			}
		}
	}
}
