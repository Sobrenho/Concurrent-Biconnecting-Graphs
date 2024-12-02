package graphs

func (graph *Graph) tarjanDFS() []int {

	visited := make([]bool, graph.VerticesCount())

	components := make([]int, 0)

	for u := 0; u < graph.VerticesCount(); u++ {
		if !visited[u] {
			graph.tarjanDFSVisit(u, visited)
			components = append(components, u)
		}
	}

	return components
}

func (graph *Graph) tarjanDFSVisit(u int, visited []bool) {

	visited[u] = true

	for _, v := range graph.Adjacents(u) {
		if !visited[v] {
			graph.tarjanDFSVisit(v, visited)
		}
	}
}

func (graph *Graph) Tarjan(vertex int) []Block {

	parent := make([]int, graph.VerticesCount())
	desc   := make([]int, graph.VerticesCount())
	ret    := make([]int, graph.VerticesCount())

	for u := 0; u < graph.VerticesCount(); u++ {
		parent[u] = -1
		desc[u] = -1
	}

	time := 0
	edgeStack := NewStack[Edge]()
	blocks := make([]Block, 0)

	graph.tarjanVisit(vertex, parent, desc, ret, &time, edgeStack, &blocks)

	if !edgeStack.IsEmpty() {

		lastBlock := make(Block, 0)

		for !edgeStack.IsEmpty() {
			lastBlock = append(lastBlock, edgeStack.Pop())
		}

		blocks = append(blocks, lastBlock)
	}

	return blocks
}

func (graph *Graph) tarjanVisit(u int, parent []int, desc []int, ret []int, time *int, edgeStack *StackX[Edge], blocks *[]Block) {

	*time++

	desc[u] = *time
	ret[u]  = *time

	childrenCount := 0

	for _, v := range graph.Adjacents(u) {

		if desc[v] == -1 {

			parent[v] = u

			childrenCount++

			edgeStack.Push(Edge{u ,v})
			graph.tarjanVisit(v, parent, desc, ret, time, edgeStack, blocks)

			ret[u] = min(ret[u], ret[v])

			if (desc[u] == 1 && childrenCount > 1) || (desc[u] > 1 && ret[v] >= desc[u]) {

				thisBlock := make(Block, 0)

				for {
					anEdge := edgeStack.Pop()
					thisBlock = append(thisBlock, anEdge)
					if anEdge.U == u && anEdge.V == v {
						break
					}
				}

				if len(thisBlock) > 0 {
					*blocks = append(*blocks, thisBlock)
				}
			}

		} else if v != parent[u] {

			ret[u] = min(ret[u], desc[v])

			if desc[v] < desc[u] {
				edgeStack.Push(Edge{u, v})
			}
		}
	}
}

func (graph *Graph) DFSTarjan() ([]int, []Block) {

	components := graph.tarjanDFS()
	blocks := make([]Block, 0, len(components))
	for _, u := range components {
		blocks = append(blocks, graph.Tarjan(u)...)
	}
	return components, blocks
}
