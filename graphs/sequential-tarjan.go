package graphs

const (
	white = iota
	gray  = iota
	black = iota
)

type edge struct {
	u int
	v int
}

func (graph *GraphX) tarjanDFS(desc []int) []int {

		color  := make([]int, graph.VerticesCount())
		parent := make([]int, graph.VerticesCount())

		for u := 0; u < graph.VerticesCount(); u++ {
			color[u] = white
			parent[u] = -1
		}

		time := 0
		componentsRepresentatives := make([]int, 0)

		for u := 0; u < graph.VerticesCount(); u++ {
			if color[u] == white {
				graph.tarjanDFSVisit(u, color, parent, desc, &time)
				componentsRepresentatives =
					append(componentsRepresentatives, u)
			}
		}

		return componentsRepresentatives
}

func (graph *GraphX) tarjanDFSVisit(u int, color []int, parent []int,
	desc []int, time *int) {

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

func (graph *GraphX) isInBiconnectedComponent(vertex int,
	desc []int) bool {

	color  := make([]int, graph.VerticesCount())
	parent := make([]int, graph.VerticesCount())
	ret    := make([]int, graph.VerticesCount())

	for u := 0; u < graph.VerticesCount(); u++ {
		color[u] = white
		parent[u] = -1
	}

	edgeStack := NewStack[edge]()
	time := 0

	return graph.isInBicconectedComponentVisit(vertex, color, parent,
		desc, ret, edgeStack, &time) == 1

}

func (graph *GraphX) isInBicconectedComponentVisit(u int, color []int,
	parent []int, desc []int, ret []int, edgeStack *StackX[edge],
	time *int) int {

		color[u] = gray
		*time += 1
		desc[u] = *time
		ret[u] = *time

		blocksCount := 0

		for _, v := range graph.Adjacents(u) {

			if color[v] == white {

				parent[v] = u
				edgeStack.Push(edge{u ,v})

				blocksCount = graph.isInBicconectedComponentVisit(v,
					color, parent, desc, ret, edgeStack, time)
					
				if ret[v] >= desc[u] {

					for {
						anEdge := edgeStack.Pop()
						if anEdge.u == u && anEdge.v == v {
							blocksCount++
							break
						}
					}

				} else {
					ret[u] = min (ret[u], ret[v])
				}
			} else if v != parent[u] {

				if desc[v] < desc[u] {
					edgeStack.Push(edge{u, v})
				}
				ret[u] = min(ret[u], desc[v])

			}
		}

		return blocksCount
}

func (graph *GraphX) Tarjan() ([]int, []int) {

	desc := make([]int, graph.VerticesCount())

	componentsRepresentatives := graph.tarjanDFS(desc)

	biconnectedComponentsRepresentatives := make([]int, 0,
		len(componentsRepresentatives))
	
	for _, u := range componentsRepresentatives {
		if graph.isInBiconnectedComponent(u, desc) {
			componentsRepresentatives =
				append(componentsRepresentatives, u)
		}
	}

	return componentsRepresentatives,
		biconnectedComponentsRepresentatives
}
