package graphs

type Graph struct {
	verticesCount int
	edgesCount    int
	adjacents     [][]int
}

func NewGraph(verticesCount int) *Graph {

	graph := new(Graph)

	graph.verticesCount = verticesCount
	graph.edgesCount = 0

	graph.adjacents = make([][]int, verticesCount)
	for i := range graph.adjacents {
		graph.adjacents[i] = make([]int, 0)
	}

	return graph
}

func (graph *Graph) AddEdge(vertexA int, vertexB int) {

	graph.adjacents[vertexA] = append(graph.adjacents[vertexA], vertexB)
	graph.adjacents[vertexB] = append(graph.adjacents[vertexB], vertexA)
	graph.edgesCount++

}

func (graph *Graph) Adjacents(vertex int) []int {
	return graph.adjacents[vertex]
}

func (graph *Graph) VerticesCount() int {
	return graph.verticesCount
}

func (graph *Graph) EdgesCount() int {
	return graph.edgesCount
}



// Gera seguinte grafo: https://i.imgur.com/8s24EVp.png
func TestGraph() *Graph{
	G := NewGraph(18)

	G.AddEdge(0,1)
	G.AddEdge(0,2)
	G.AddEdge(1,3)
	G.AddEdge(2,3)
	G.AddEdge(2,4)
	G.AddEdge(3,4)
	G.AddEdge(4,5)
	G.AddEdge(6,7)
	G.AddEdge(6,8)
	G.AddEdge(7,8)
	G.AddEdge(9,10)
	G.AddEdge(9,11)
	G.AddEdge(10,11)
	G.AddEdge(10,12)
	G.AddEdge(11,12)
	G.AddEdge(13,14)
	G.AddEdge(13,16)
	G.AddEdge(14,15)
	G.AddEdge(15,16)

	return G
}