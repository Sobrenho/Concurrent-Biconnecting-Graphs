package scripts

import (
	//"fmt"
	"log"
	"os"
	"splatoon-tarjan-demo/graphs"
	"time"
)

func RunDFSTarjan(args []string) int64{

	inputFile, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
		return -1
	}
	defer inputFile.Close()

	outputFile, err := os.Create(args[1])
	if err != nil {
		log.Fatal(err)
		return -1
	}
	defer outputFile.Close()

	graph, err := graphs.ReadGraphFrom(inputFile)
	if err != nil {
		log.Fatal(err)
		return -1 
	}

	beforeAlgorithm := time.Now().UnixMilli()
	_, blocks := graph.DFSTarjan()
	afterAlgorithm := time.Now().UnixMilli()

	//fmt.Println(afterAlgorithm - beforeAlgorithm)

	sortSlicesOfBlocks(blocks)
	writeSliceOfBlocksToFile(blocks, outputFile)

	return afterAlgorithm - beforeAlgorithm
}