package scripts

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
	"trabfinal/graphs"
)

func compareEdges(a graphs.Edge, b graphs.Edge) int {
	if a.U == b.U {
		return a.V - b.V
	}
	return a.U - b.U
}

func compareSlicesOfEdges(a []graphs.Edge, b []graphs.Edge) int {
	for i := range a {
		compare := compareEdges(a[i], b[i])
		if compare != 0 {
			return compare
		}
	}
	return 0
}

func compareSlicesOfSlicesOfEdges(a [][]graphs.Edge, b [][]graphs.Edge) int {
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

func sortSliceOfEdges(slice []graphs.Edge) {
	for i := range slice {
		sortEdge(&slice[i])
	}
	sort.Slice(slice, func(i int, j int) bool {
		return compareEdges(slice[i], slice[j]) < 0
	})
}

func sortSliceOfSlicesOfEdges(slice [][]graphs.Edge) {
	for _, subslice := range slice {
		sortSliceOfEdges(subslice)
	}
	sort.Slice(slice, func(i int, j int) bool {
		return compareSlicesOfEdges(slice[i], slice[j]) < 0
	})
}

func writeEdgeToFile(edge graphs.Edge, file *os.File) error {

	err := binary.Write(file, binary.BigEndian, int64(edge.U))
	if err != nil {
		return err
	}

	err = binary.Write(file, binary.BigEndian, int64(edge.V))
	if err != nil {
		return err
	}

	return nil
}

func writeSliceOfEdgesToFile(slice []graphs.Edge, file *os.File) error {

	err := binary.Write(file, binary.BigEndian, int64(len(slice)))
	if err != nil {
		return err
	}

	for _, item := range slice {
		err = writeEdgeToFile(item, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeSliceOfSlicesOfEdgesToFile(slice [][]graphs.Edge, file *os.File) error {

	err := binary.Write(file, binary.BigEndian, int64(len(slice)))
	if err != nil {
		return err
	}

	for _, subslice := range slice {
		err = writeSliceOfEdgesToFile(subslice, file)
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

	beforeAlgorithm := time.Now().UnixMilli()
	_, blocks := graph.DFSTarjan()
	afterAlgorithm := time.Now().UnixMilli()

	fmt.Println(afterAlgorithm - beforeAlgorithm)

	sortSliceOfSlicesOfEdges(blocks)
	writeSliceOfSlicesOfEdgesToFile(blocks, outputFile)
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

	beforeAlgorithm := time.Now().UnixMilli()
	_, blocks := graph.SplatoonTarjan(threadsCount)
	afterAlgorithm := time.Now().UnixMilli()

	fmt.Println(afterAlgorithm - beforeAlgorithm)

	sortSliceOfSlicesOfEdges(blocks)
	writeSliceOfSlicesOfEdgesToFile(blocks, outputFile)
}

func ValidateSplatoonTarjan(args []string) {

	type graphSizeSetting struct {
		verticesCount int
		edgesCount int
	}

	iterationsPerGraphSize, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
		return
	}

	graphSizesToTest := []graphSizeSetting {
		{1000,  2000},
		{1000,  3000},
		{1000,  4000},
		{2000,  4000},
		{2000,  8000},
		{2000, 10000},
		{5000,  7000},
		{5000, 10000},
		{5000, 13000},
	}

	threadsCountsToTest := []int {1, 2, 4, 8}

	totalTests := len(graphSizesToTest) * iterationsPerGraphSize * len(threadsCountsToTest)
	rightAnswers := 0
	testsDone := 0

	for _, graphSize := range graphSizesToTest {

		for i := 0; i < iterationsPerGraphSize; i++ {

			graph := graphs.NewRandomGraph(graphSize.verticesCount, graphSize.edgesCount)

			_, blocksDFSTarjan := graph.DFSTarjan()
			sortSliceOfSlicesOfEdges(blocksDFSTarjan)

			for _, threadsCount := range threadsCountsToTest {

				_, blocksSplatoonTarjan := graph.SplatoonTarjan(threadsCount)
				sortSliceOfSlicesOfEdges(blocksSplatoonTarjan)

				if compareSlicesOfSlicesOfEdges(blocksDFSTarjan, blocksSplatoonTarjan) == 0 {
					rightAnswers++
				}

				testsDone++

				fmt.Printf("%d/%d tests done.\n", testsDone, totalTests)
			}
		}
	}

	fmt.Printf("%d/%d right.", rightAnswers, totalTests)
}