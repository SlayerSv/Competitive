package graphs

var n int
var graph = make([][]int, n)
var visited = make([]bool, n)
var tin = make([]int, n)
var lowest = make([]int, n)
var bridges = [][2]int{}
var time = 0

func findBridges(node, parent int) {
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
			findBridges(neigh, node)
			lowest[node] = min(lowest[node], lowest[neigh])
			if lowest[neigh] > tin[node] {
				bridges = append(bridges, [2]int{node, neigh})
			}
		}
	}
}
