package graphs

const (
	white = iota
	gray  = iota
	black = iota
)

type Edge struct {
	U int
	V int
}

func (graph *GraphX) tarjanDFS() []int {

	color  := make([]int, graph.VerticesCount())
	parent := make([]int, graph.VerticesCount())
	desc   := make([]int, graph.VerticesCount())

	for u := 0; u < graph.VerticesCount(); u++ {
		color[u] = white
		parent[u] = -1
	}

	time := 0
	components := make([]int, 0)

	for u := 0; u < graph.VerticesCount(); u++ {
		if color[u] == white {
			graph.tarjanDFSVisit(u, color, parent, desc, &time)
			components = append(components, u)
		}
	}

	return components
}

func (graph *GraphX) tarjanDFSVisit(u int, color []int, parent []int, desc []int, time *int) {

	*time += 1
	desc[u] = *time
	color[u] = gray
	for _, v := range graph.Adjacents(u) {
		if color[v] == white {
			parent[v] = u
			graph.tarjanDFSVisit(v, color, parent, desc, time)
		}
	}
	color[u] = black
	*time += 1
}

func (graph *GraphX) Tarjan(vertex int) [][]Edge {

	parent := make([]int, graph.VerticesCount())
	desc   := make([]int, graph.VerticesCount())
	ret    := make([]int, graph.VerticesCount())

	for u := 0; u < graph.VerticesCount(); u++ {
		parent[u] = -1
		desc[u] = -1
	}

	time := 0
	edgeStack := NewStack[Edge]()
	blocks := make([][]Edge, 0)

	graph.tarjanVisit(vertex, parent, desc, ret, &time, edgeStack, &blocks)

	if !edgeStack.IsEmpty() {

		lastBlock := make([]Edge, 0)

		for !edgeStack.IsEmpty() {
			lastBlock = append(lastBlock, edgeStack.Pop())
		}

		blocks = append(blocks, lastBlock)
	}

	return blocks
}

func (graph *GraphX) tarjanVisit(u int, parent []int, desc []int, ret []int, time *int, edgeStack *StackX[Edge], blocks *[][]Edge) {

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

				thisBlock := make([]Edge, 0)

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

func (graph *GraphX) DFSTarjan() ([]int, [][]Edge) {

	components := graph.tarjanDFS()
	blocks := make([][]Edge, 0, len(components))
	for _, u := range components {
		blocks = append(blocks, graph.Tarjan(u)...)
	}
	return components, blocks
}
