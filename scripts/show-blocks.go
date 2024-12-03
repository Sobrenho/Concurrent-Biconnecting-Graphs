package scripts

import (
	"fmt"
	"log"
	"os"
	"splatoon-tarjan-demo/graphs"
)

func ShowBlocks(args []string) {

	inputFile, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer inputFile.Close()

	var blocks [][]graphs.Edge
	err = readSliceOfSlicesOfEdgesFromFile(&blocks, inputFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(blocks)
}