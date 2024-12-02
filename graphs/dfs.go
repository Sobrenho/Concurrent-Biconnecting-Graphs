package graphs

func (graph *Graph) DFSComponentSearch() []int {

	visited := make([]bool, graph.VerticesCount())

	var visit func(int)
	visit = func(u int) {

		visited[u] = true

		for _, v := range graph.Adjacents(u) {
			if !visited[v] {
				visit(v)
			}
		}
	}

	components := make([]int, 0)

	for u := 0; u < graph.VerticesCount(); u++ {
		if !visited[u] {
			visit(u)
			components = append(components, u)
		}
	}

	return components
}