package scripts

import (
	"log"
	"os"
	"splatoon-tarjan-demo/graphs"
	"strconv"
	"time"
)

func RunShiloachVishkinTarjan(args []string) int64{

	threadsCount, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
		return -1
	}

	inputFile, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
		return -1 
	}
	defer inputFile.Close()

	outputFile, err := os.Create(args[2])
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
	_, blocks := graph.ShiloachVishkinTarjan(threadsCount)
	afterAlgorithm := time.Now().UnixMilli()

	//fmt.Println(afterAlgorithm - beforeAlgorithm)

	sortSlicesOfBlocks(blocks)
	writeSliceOfBlocksToFile(blocks, outputFile)

	return afterAlgorithm - beforeAlgorithm
}