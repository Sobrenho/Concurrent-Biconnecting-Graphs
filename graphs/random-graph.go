package graphs

import "math/rand/v2"

func NewRandomGraph(verticesCount int, edgesCount int) *Graph {

	graph := NewGraph(verticesCount)

	edgeExists := make(map[int]map[int]bool, verticesCount)
	for u := 0; u < verticesCount; u++ {
		edgeExists[u] = make(map[int]bool, verticesCount)
	}

	for graph.EdgesCount() < edgesCount {

		u := rand.IntN(verticesCount)
		v := rand.IntN(verticesCount)

		if edgeExists[u][v] {
			continue
		}

		graph.AddEdge(u, v)
		edgeExists[u][v] = true
		edgeExists[v][u] = true
	}

	return graph

}