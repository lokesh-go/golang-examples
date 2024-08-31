package graph

import "fmt"

// DFS performs depth-first search on a graph represented by an adjacency list.
func DFS(graph map[int][]int, start int, visited map[int]bool) {
	// Mark the current node as visited
	visited[start] = true
	fmt.Printf("%d ", start)

	// Recursively visit all the adjacent nodes that have not been visited yet
	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			DFS(graph, neighbor, visited)
		}
	}
}

/*
func main() {
	// Example graph represented as an adjacency list
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {},
		6: {},
	}

	// Map to track visited nodes
	visited := make(map[int]bool)

	// Start DFS from node 1
	DFS(graph, 1, visited)
}
*/
