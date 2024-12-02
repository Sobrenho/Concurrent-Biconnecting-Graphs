package scripts

import (
	"fmt"
	"log"
	"os"
	"time"
	"trabfinal/graphs"
)

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

	sortSlicesOfBlocks(blocks)
	writeSliceOfBlocksToFile(blocks, outputFile)
}