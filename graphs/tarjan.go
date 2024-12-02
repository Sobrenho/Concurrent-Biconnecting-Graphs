package graphs

import "trabfinal/structures"

func (graph *Graph) Tarjan(vertex int) []Block {

	parent := make([]int, graph.VerticesCount())
	desc   := make([]int, graph.VerticesCount())
	ret    := make([]int, graph.VerticesCount())

	for u := 0; u < graph.VerticesCount(); u++ {
		parent[u] = -1
		desc[u] = -1
	}

	time := 0
	edgeStack := structures.MakeStack[Edge]()
	blocks := make([]Block, 0)

	var visit func(int)
	visit = func(u int) {

		time++

		desc[u] = time
		ret[u]  = time

		childrenCount := 0

		for _, v := range graph.Adjacents(u) {

			if desc[v] == -1 {

				parent[v] = u

				childrenCount++

				edgeStack.Push(Edge{u ,v})
				visit(v)

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
						blocks = append(blocks, thisBlock)
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

	visit(vertex)

	if !edgeStack.IsEmpty() {

		lastBlock := make(Block, 0)

		for !edgeStack.IsEmpty() {
			lastBlock = append(lastBlock, edgeStack.Pop())
		}

		blocks = append(blocks, lastBlock)
	}

	return blocks
}
