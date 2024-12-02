package tests

import (
	"fmt"
	"sort"
	"trabfinal/graphs"
)

func compareEdges(a graphs.Edge, b graphs.Edge) int {
	if a.U == b.U {
		return a.V - b.V
	}
	return a.U - b.U
}

func compareSlicesOfEdges(a graphs.Block, b graphs.Block) int {
	for i := range a {
		compare := compareEdges(a[i], b[i])
		if compare != 0 {
			return compare
		}
	}
	return 0
}

func compareSlicesOfSlicesOfEdges(a []graphs.Block, b []graphs.Block) int {
	for i := range a {
		compare := compareSlicesOfEdges(a[i], b[i])
		if compare != 0 {
			return compare
		}
	}
	return 0
}

func sortEdge(edge *graphs.Edge) {
	if edge.U > edge.V {
		edge.U, edge.V = edge.V, edge.U
	}
}

func sortSliceOfEdges(slice graphs.Block) {
	for i := range slice {
		sortEdge(&slice[i])
	}
	sort.Slice(slice, func(i int, j int) bool {
		return compareEdges(slice[i], slice[j]) < 0
	})
}

func sortSliceOfSlicesOfEdges(slice []graphs.Block) {
	for _, subslice := range slice {
		sortSliceOfEdges(subslice)
	}
	sort.Slice(slice, func(i int, j int) bool {
		return compareSlicesOfEdges(slice[i], slice[j]) < 0
	})
}

func ValidateSplatoonTarjan(iterations int) {

	wrongs := 0

	for i := 0; i < iterations; i++ {

		graph := graphs.NewRandomGraph(1000, 2000)

		_, blocksDFS := graph.DFSTarjan()
		sortSliceOfSlicesOfEdges(blocksDFS)

		for _, t := range []int{1, 2, 4, 8} {

			_, blocksSplatoon := graph.SplatoonTarjan(t)
			sortSliceOfSlicesOfEdges(blocksSplatoon)

			if compareSlicesOfSlicesOfEdges(blocksDFS, blocksSplatoon) == 0 {
				fmt.Println("OK!")
			} else {
				fmt.Println("Wrong!")
				wrongs++
			}

		}
	}

	fmt.Println(wrongs)
}