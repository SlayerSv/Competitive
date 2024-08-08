package graphs

func findBridges(graph [][]int) [][2]int {
	n := len(graph)
	visited := make([]bool, n)
	lowest := make([]int, n)
	bridges := [][2]int{}
	tin := make([]int, n)
	time := 0
	var dfs func(int, int)
	dfs = func(node, parent int) {
		visited[node] = true
		tin[node] = time
		lowest[node] = time
		time++
		for _, neigh := range graph[node] {
			if parent == neigh {
				continue
			}
			if visited[neigh] {
				lowest[node] = min(lowest[node], tin[neigh])
			} else {
				dfs(neigh, node)
				lowest[node] = min(lowest[node], lowest[neigh])
				if lowest[neigh] > tin[node] {
					bridges = append(bridges, [2]int{node, neigh})
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, i)
		}
	}
	return bridges
}
