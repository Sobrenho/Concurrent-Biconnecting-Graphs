package scripts

import (
	"encoding/binary"
	"log"
	"os"
	"sort"
	"strconv"
	"trabfinal/graphs"
)

func compareSlices(a []int, b []int) int {
	for i := range a {
		if a[i] != b[i] {
			return a[i] - b[i]
		}
	}
	return 0
}

func sortSlice(slice []int) {
	sort.Slice(slice, func(i int, j int) bool {
		return slice[i] < slice[j]
	})
}

func sortSliceOfSlices(slice [][]int) {
	for _, slice := range slice {
		sortSlice(slice)
	}
	sort.Slice(slice, func(i int, j int) bool {
		return compareSlices(slice[i], slice[j]) < 0
	})
}

func verticesOfComponentDFS(graph *graphs.GraphX, u int, visited []bool, vertices *[]int) {
	visited[u] = true
	*vertices = append(*vertices, u)
	for _, v := range graph.Adjacents(u) {
		if !visited[v] {
			verticesOfComponentDFS(graph, v, visited, vertices)
		}
	}
}

func verticesOfComponent(graph *graphs.GraphX, vertex int) []int {
	visited := make([]bool, graph.VerticesCount())
	vertices := make([]int, 0)
	verticesOfComponentDFS(graph, vertex, visited, &vertices)
	return vertices
}

func writeSliceToFile(slice []int, file *os.File) error {

	err := binary.Write(file, binary.BigEndian, int64(len(slice)))
	if err != nil {
		return err
	}

	for _, item := range slice {
		err = binary.Write(file, binary.BigEndian, int64(item))
		if err != nil {
			return err
		}
	}

	return nil
}

func writeSliceOfSlicesToFile(slice [][]int, file *os.File) error {

	err := binary.Write(file, binary.BigEndian, int64(len(slice)))
	if err != nil {
		return err
	}

	for _, subslice := range slice {
		err = writeSliceToFile(subslice, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func MakeRandomGraphFile(args []string) {
	
	file, err := os.Create(args[2])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	verticesCount, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
		return
	}

	edgesCount, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	graph := graphs.NewRandomGraph(verticesCount, edgesCount)
	err = graph.WriteTo(file)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func RunDFSTarjan(args []string) {

	inputFile, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(args[1])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer outputFile.Close()

	graph, err := graphs.ReadGraphFrom(inputFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, components := graph.DFSTarjan()

	componentsExpanded := make([][]int, len(components))
	for i, component := range components {
		componentsExpanded[i] = verticesOfComponent(graph, component)
	}

	sortSliceOfSlices(componentsExpanded)
	writeSliceOfSlicesToFile(componentsExpanded, outputFile)
}

func RunSplatoonTarjan(args []string) {

	threadsCount, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	inputFile, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(args[2])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer outputFile.Close()

	graph, err := graphs.ReadGraphFrom(inputFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, components := graph.SplatoonTarjan(threadsCount)

	componentsExpanded := make([][]int, len(components))
	for i, component := range components {
		componentsExpanded[i] = verticesOfComponent(graph, component)
	}

	sortSliceOfSlices(componentsExpanded)
	writeSliceOfSlicesToFile(componentsExpanded, outputFile)
}
