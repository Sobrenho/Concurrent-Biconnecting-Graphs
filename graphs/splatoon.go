package graphs

func SplatoonComponentSearch(graph Graph, threadsCount int) []*Vertex {

	// Union-find 

	verticesChannel := make(chan *Vertex)
	go func() {
		for i := 0; i < graph.N; i++ {
			verticesChannel <- &graph.Vertices[i]
		}
	}()

	canFinish := make(chan bool)

	verticesConsumed := make(chan int)
	verticesConsumed <- 0

	isVisited := make(map[*Vertex]bool, graph.N)

	for i := 0; i < threadsCount; i++ {

		go func(verticesChannel chan *Vertex, canFinish chan bool, verticesConsumed chan int) {

			for {
				select {
	
				case vertex := <- verticesChannel:
	
					isVisited[vertex] = true
	
					/*
					for _, neighbor := range graph.AdjancencyList[vertex] {
	
						// join (vertex, neighbor)
	
					}
					*/
	
					vConsumed := <- verticesConsumed
					
					vConsumed++
	
					verticesConsumed <- vConsumed
	
					if vConsumed == graph.N {
						canFinish <- true
						return
					}
	
				case <- canFinish:
	
					canFinish <- true
					return
				
				}
	
			}
	
		}(verticesChannel, canFinish, verticesConsumed)

	}

	// Get from union-find
	return make([]*Vertex, 10)

}