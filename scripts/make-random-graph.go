package scripts

import (
	"log"
	"os"
	"strconv"
	"trabfinal/graphs"
)

func MakeRandomGraph(args []string) {
	
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