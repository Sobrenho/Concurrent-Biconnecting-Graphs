package scripts

import (
	"fmt"
	"log"
	"os"
	"splatoon-tarjan-demo/graphs"
)

func ShowGraph(args []string) {

	inputFile, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer inputFile.Close()

	graph, err := graphs.ReadGraphFrom(inputFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	vertices := make([]int, graph.VerticesCount())
	for u := 0; u < graph.VerticesCount(); u++ {
		vertices[u] = u
	}

	fmt.Print("Vertices: ")
	fmt.Println(vertices)
	fmt.Println("Adjacency lists:")
	for u := 0; u < graph.VerticesCount(); u++ {
		fmt.Printf("  %d: ", u)
		fmt.Println(graph.Adjacents(u))
	}
}