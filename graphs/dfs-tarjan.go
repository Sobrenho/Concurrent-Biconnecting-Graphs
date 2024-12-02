package graphs

func (graph *Graph) DFSTarjan() ([]int, []Block) {

	components := graph.DFSComponentSearch()
	blocks := make([]Block, 0, len(components))
	for _, u := range components {
		blocks = append(blocks, graph.Tarjan(u)...)
	}
	return components, blocks
}
