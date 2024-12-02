package scripts

import "splatoon-tarjan-demo/graphs"

func compareEdges(a graphs.Edge, b graphs.Edge) int {
	if a.U == b.U {
		return a.V - b.V
	}
	return a.U - b.U
}

func compareBlocks(a graphs.Block, b graphs.Block) int {
	for i := range a {
		compare := compareEdges(a[i], b[i])
		if compare != 0 {
			return compare
		}
	}
	return 0
}

func compareSlicesOfBlocks(a []graphs.Block, b []graphs.Block) int {
	for i := range a {
		compare := compareBlocks(a[i], b[i])
		if compare != 0 {
			return compare
		}
	}
	return 0
}