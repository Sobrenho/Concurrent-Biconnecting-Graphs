package graphs

type GraphX struct {
	verticesCount int
	adjacents     [][]int
}

func NewGraphX(verticesCount int) *GraphX {

	graph := new(GraphX)

	graph.verticesCount = verticesCount

	graph.adjacents = make([][]int, verticesCount)
	for i := range graph.adjacents {
		graph.adjacents[i] = make([]int, 0)
	}

	return graph
}

func (graph *GraphX) AddEdge(vertexA int, vertexB int) {

	graph.adjacents[vertexA] = append(graph.adjacents[vertexA], vertexB)
	graph.adjacents[vertexB] = append(graph.adjacents[vertexB], vertexA)

}

func (graph *GraphX) Adjacents(vertex int) []int {
	return graph.adjacents[vertex]
}
