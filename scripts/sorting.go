package scripts

import (
	"sort"
	"trabfinal/graphs"
)

func sortEdge(edge *graphs.Edge) {
	if edge.U > edge.V {
		edge.U, edge.V = edge.V, edge.U
	}
}

func sortBlock(block graphs.Block) {
	for i := range block {
		sortEdge(&block[i])
	}
	sort.Slice(block, func(i int, j int) bool {
		return compareEdges(block[i], block[j]) < 0
	})
}

func sortSlicesOfBlocks(slice []graphs.Block) {
	for _, subslice := range slice {
		sortBlock(subslice)
	}
	sort.Slice(slice, func(i int, j int) bool {
		return compareBlocks(slice[i], slice[j]) < 0
	})
}