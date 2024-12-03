package scripts

import (
	"fmt"
	"log"
	"splatoon-tarjan-demo/graphs"
	"strconv"
)

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

	threadsCountsToTest := []int {1, 2, 4, 8, 16}

	totalTests := len(graphSizesToTest) * iterationsPerGraphSize * len(threadsCountsToTest)
	rightAnswers := 0
	testsDone := 0

	for _, graphSize := range graphSizesToTest {

		for i := 0; i < iterationsPerGraphSize; i++ {

			graph := graphs.NewRandomGraph(graphSize.verticesCount, graphSize.edgesCount)

			_, blocksDFSTarjan := graph.DFSTarjan()
			sortSlicesOfBlocks(blocksDFSTarjan)

			for _, threadsCount := range threadsCountsToTest {

				_, blocksSplatoonTarjan := graph.SplatoonTarjan(threadsCount)
				sortSlicesOfBlocks(blocksSplatoonTarjan)

				if compareSlicesOfBlocks(blocksDFSTarjan, blocksSplatoonTarjan) == 0 {
					rightAnswers++
				}

				testsDone++

				fmt.Printf("%d/%d tests done.\n", testsDone, totalTests)
			}
		}
	}

	fmt.Printf("%d/%d right.\n", rightAnswers, totalTests)
}