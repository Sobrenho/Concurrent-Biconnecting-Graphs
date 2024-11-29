package tests

import (
	"fmt"
	"reflect"
	"sort"
	"trabfinal/graphs"
)

func getVerticesOfComponentDFS(graph *graphs.GraphX, u int, visited []bool, visitedVertices *[]int) {

	visited[u] = true
	*visitedVertices = append(*visitedVertices, u)
	for _, v := range graph.Adjacents(u) {
		if !visited[v] {
			getVerticesOfComponentDFS(graph, v, visited, visitedVertices)
		}
	}

}

func getVerticesOfComponent(graph *graphs.GraphX, source int) []int {

	visited := make([]bool, graph.VerticesCount())
	visitedVertices := make([]int, 0)
	getVerticesOfComponentDFS(graph, source, visited, &visitedVertices)
	return visitedVertices

}

func compareSlices(a []int, b []int) int {
	for i := range a {
		if a[i] != b[i] {
			return a[i] - b[i]
		}
	}
	return 0
}

func sortSliceOfSlices(slice [][]int) {
	
	sort.Slice(slice, func (i int, j int) bool {
		return compareSlices(slice[i], slice[j]) < 0
	})

}

func ValidateSplatoonTarjan(iterations int) int {

	wrongs := 0

	for i := 0; i < iterations; i++ {

		graph := graphs.NewRandomGraph(1000, 2000)

		_, blockRepsDFSTarjan := graph.DFSTarjan()
		blocksDFSTarjan := make([][]int, len(blockRepsDFSTarjan))
		for i, blockRep := range blockRepsDFSTarjan {
			blocksDFSTarjan[i] = getVerticesOfComponent(graph, blockRep)
		}
		sortSliceOfSlices(blocksDFSTarjan)

		for _, t := range []int{1, 2, 4, 8} {

			_, blockRepsSplatoonTarjan := graph.SplatoonTarjan(t)
			blocksSplatoonTarjan := make([][]int, len(blockRepsSplatoonTarjan))
			for i, blockRep := range blockRepsSplatoonTarjan {
				blocksSplatoonTarjan[i] = getVerticesOfComponent(graph, blockRep)
			}
			sortSliceOfSlices(blocksSplatoonTarjan)

			if reflect.DeepEqual(blocksDFSTarjan, blocksSplatoonTarjan) {
				fmt.Println("OK!")
			} else {
				fmt.Println("Wrong!")
				wrongs++
			}

		}
	}

	return wrongs
}